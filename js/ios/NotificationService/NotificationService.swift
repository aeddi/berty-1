//
//  NotificationService.swift
//  NotificationService
//
//  Created by Antoine Eddi on 7/24/21.
//

import Bertypush
import UserNotifications
import OSLog

class NotificationService: UNNotificationServiceExtension {
  let oslogger = OSLog(subsystem: "tech.berty.ios.notif", category: "notif")

  var contentHandler: ((UNNotificationContent) -> Void)?
  var bestAttemptContent: UNMutableNotificationContent?


  override func didReceive(_ request: UNNotificationRequest, withContentHandler contentHandler: @escaping (UNNotificationContent) -> Void) {
    os_log("didReceive encrypted notification", log: oslogger)

    self.contentHandler = contentHandler
    self.bestAttemptContent = (request.content.mutableCopy() as? UNMutableNotificationContent)

    var rootDir: String
    do {
      rootDir = try RootDirGet()
    } catch {
      os_log("getting root dir failed: %@", log: oslogger, type: .error, error.localizedDescription)
      displayFallback()
      return
    }

    guard let data = request.content.userInfo["data"] as? String else {
      os_log("No filed 'data' found in received push notification", log: oslogger, type: .error)
      displayFallback()
      return
    }

    os_log("WIP_LOG starting push decrypt", log: oslogger)

    var error: NSError?
    guard let decrypted = BertypushPushDecryptStandalone(rootDir, data, &error) else {
       if let error = error {
          os_log("Push notif decryption failed: %{public}@", log: oslogger, error.localizedDescription)
       } else {
          os_log("Push notif decryption failed without returning error", log: oslogger)
       }
       displayFallback()
       return
    }

    os_log("WIP_LOG decrypt end: %@", log: oslogger, decrypted.content)

    // TODO: implement display logic
    self.bestAttemptContent!.title = "Decrypted"
    self.bestAttemptContent!.subtitle = request.content.userInfo["data"] as! String
    self.bestAttemptContent!.body = decrypted.content
    // self.bestAttemptContent!.body = "\(decrypted.accountName) \(decrypted.accountID) \(decrypted.memberDisplayName) \(decrypted.payloadAttrsJSON)"
    contentHandler(self.bestAttemptContent!)


  }

  // This callback will be called by iOS if data decryption took to much time
  override func serviceExtensionTimeWillExpire() {
    os_log("push notif decryption timed out", log: oslogger, type: .error)
    displayFallback()
  }

  func displayFallback() {
    // TODO: add i18n
    self.bestAttemptContent!.title = "Fallback Title"
    self.bestAttemptContent!.subtitle = "Fallback Subtitle"
    self.bestAttemptContent!.body = "Fallback Body"

    // Display fallback content
    self.contentHandler!(self.bestAttemptContent!)
  }
}
