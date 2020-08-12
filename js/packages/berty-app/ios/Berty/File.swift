//
//  File.swift
//  Berty Debug
//
//  Created by Guilhem Fanton on 08/01/2020.
//  Copyright © 2020 Berty Technologies. All rights reserved.
//

import Foundation
import go_bridge

class Bridge: NSObject {
  static func setBridgeBackgroundID(bid: String) {
    LifeCycleDriver.BackgroundID = bid
  }
}
