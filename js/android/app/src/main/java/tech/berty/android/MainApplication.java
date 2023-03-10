package tech.berty.android;

import android.app.Application;

import androidx.lifecycle.Lifecycle;
import androidx.lifecycle.LifecycleObserver;
import androidx.lifecycle.OnLifecycleEvent;
import androidx.lifecycle.ProcessLifecycleOwner;

import com.facebook.react.PackageList;
import com.facebook.react.ReactApplication;
import com.facebook.react.ReactNativeHost;
import com.facebook.react.ReactPackage;
import com.facebook.react.defaults.DefaultNewArchitectureEntryPoint;
import com.facebook.react.defaults.DefaultReactNativeHost;
import com.facebook.soloader.SoLoader;

import android.content.res.Configuration;
import expo.modules.ApplicationLifecycleDispatcher;
import expo.modules.ReactNativeHostWrapper;

import java.util.List;

import tech.berty.addressbook.AddressBookPackage;
import tech.berty.gobridge.GoBridgePackage;
import tech.berty.notification.NotificationPackage;
import tech.berty.notification.NotificationService;
import tech.berty.rootdir.RootDirPackage;
import tech.berty.gobridge.Logger;

public class MainApplication extends Application implements ReactApplication, LifecycleObserver {
    private static final String TAG = "MainApplication";

    public enum AppState {
        Background,
        Foreground
    }

    private static AppState appState = AppState.Foreground;

    private final ReactNativeHost mReactNativeHost =
        new ReactNativeHostWrapper(this, new DefaultReactNativeHost(this) {
            @Override
            public boolean getUseDeveloperSupport() {
                return BuildConfig.DEBUG;
            }

            @Override
            protected List<ReactPackage> getPackages() {
                List<ReactPackage> packages = new PackageList(this).getPackages();
                // Packages that cannot be autolinked yet can be added manually here, for example:
                // packages.add(new MyReactNativePackage());
                packages.add(new NotificationPackage());
                packages.add(new RootDirPackage());
                packages.add(new GoBridgePackage());
                packages.add(new AddressBookPackage());
                return packages;
            }

            @Override
            protected String getJSMainModuleName() {
                return "index";
            }
            
            @Override
            protected boolean isNewArchEnabled() {
              return BuildConfig.IS_NEW_ARCHITECTURE_ENABLED;
            }
            
            @Override
            protected Boolean isHermesEnabled() {
              return BuildConfig.IS_HERMES_ENABLED;
            }
        });



    @OnLifecycleEvent(Lifecycle.Event.ON_STOP)
    public void onAppBackgrounded() {
        //App in background
        MainApplication.appState = AppState.Background;
        Logger.d(TAG, "AppState" + MainApplication.appState);
    }

    @OnLifecycleEvent(Lifecycle.Event.ON_START)
    public void onAppForegrounded() {
        //App in foreground
        MainApplication.appState = AppState.Foreground;
        Logger.d(TAG, "AppState" + MainApplication.appState);
    }

    @Override
    public ReactNativeHost getReactNativeHost() {
        return mReactNativeHost;
    }

    @Override
    public void onCreate() {
        super.onCreate();

        // init SoLoader
        SoLoader.init(this, /* native exopackage */ false);

        if (BuildConfig.IS_NEW_ARCHITECTURE_ENABLED) {
            // If you opted-in for the New Architecture, we load the native entry point for this app.
            DefaultNewArchitectureEntryPoint.load();
        }
        ReactNativeFlipper.initializeFlipper(this, getReactNativeHost().getReactInstanceManager());

        // register for lifecycle events
        ProcessLifecycleOwner.get().getLifecycle().addObserver(this);

        // init expo
        ApplicationLifecycleDispatcher.onApplicationCreate(this);
    }

    @Override
    public void onConfigurationChanged(Configuration newConfig) {
        super.onConfigurationChanged(newConfig);
        ApplicationLifecycleDispatcher.onConfigurationChanged(this, newConfig);
    }
}
