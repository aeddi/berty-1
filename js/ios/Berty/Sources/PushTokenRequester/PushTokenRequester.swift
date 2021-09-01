//
//  PushTokenRequester.swift
//  Berty
//
//  Created by Antoine Eddi on 02/08/2021.
//

import UserNotifications

@objc(PushTokenRequester)
class PushTokenRequester: NSObject {
  @objc static var shared: PushTokenRequester = PushTokenRequester()
  static let requestSema = DispatchSemaphore(value: 1)
  static let responseSema = DispatchSemaphore(value: 0)
  static var error: NSError?
  static var token: NSData?

  @objc func request(_ resolve: @escaping RCTPromiseResolveBlock, reject: @escaping RCTPromiseRejectBlock) {
    PushTokenRequester.requestSema.wait()
    defer {
      PushTokenRequester.error = nil
      PushTokenRequester.token = nil
      PushTokenRequester.requestSema.signal()
    }
    
    UNUserNotificationCenter.current().getNotificationSettings { settings in
      switch settings.authorizationStatus {
      case .authorized, .provisional:
        DispatchQueue.main.async {
          UIApplication.shared.registerForRemoteNotifications()
        }
        PushTokenRequester.responseSema.wait()
        
        if let token = PushTokenRequester.token,
           let bundleID = Bundle.main.bundleIdentifier {
          if let jsonData = try? JSONSerialization.data(withJSONObject: [
            "token": token.base64EncodedString(options: NSData.Base64EncodingOptions(rawValue: 0)),
            "bundleId": bundleID,
          ], options: []) {
            resolve(String(data: jsonData, encoding: .ascii))
            return
          }

          PushTokenRequester.error = NSError(domain: "push", code: 1, userInfo: [NSLocalizedDescriptionKey: "can't serialize token request response"])
        }

      default:
        PushTokenRequester.error = NSError(domain: "push", code: 1, userInfo: [NSLocalizedDescriptionKey: "notification permission not granted"])
      }
      
      var error: Error
      if PushTokenRequester.error != nil {
        error = PushTokenRequester.error!
      } else {
        error = NSError(domain: "push", code: 1, userInfo: [NSLocalizedDescriptionKey: "request failed for unknown reason"])
      }
      
      reject("token_request_failed", error.localizedDescription, error)
    }
  }

  @objc func onRequestSucceeded(_ deviceToken: NSData) {
    PushTokenRequester.token = deviceToken
    PushTokenRequester.responseSema.signal()
  }

  @objc func onRequestFailed(_ requestError: NSError) {
    PushTokenRequester.error = requestError
    PushTokenRequester.responseSema.signal()
  }
  
  @objc static func requiresMainQueueSetup() -> Bool {
      return false
  }
}
