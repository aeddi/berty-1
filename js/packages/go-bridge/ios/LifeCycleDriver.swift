//
//  LifeCycleDriver.swift
//  GoBridge
//
//  Created by Guilhem Fanton on 11/08/2020.
//  Copyright © 2020 Berty Technologies. All rights reserved.
//

import Foundation
import BackgroundTasks
import Bertybridge

public class LifeCycleDriver: NSObject {
    public static var Shared: LifeCycleDriver = LifeCycleDriver()
    public static var BackgroundID: String = ""
    let handlers: [BertybridgeAppStateHandlerProtocol] = []
    
    
}
