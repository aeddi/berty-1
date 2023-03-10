package tech.berty.android;

import android.os.Bundle; // needed by react-native-bootsplash

import expo.modules.ReactActivityDelegateWrapper;
import com.facebook.react.ReactActivity;
import com.facebook.react.ReactActivityDelegate;
import com.facebook.react.defaults.DefaultNewArchitectureEntryPoint;
import com.facebook.react.defaults.DefaultReactActivityDelegate;
import com.zoontek.rnbootsplash.RNBootSplash; // needed by react-native-bootsplash
import com.facebook.react.ReactActivityDelegate;
import com.facebook.react.ReactRootView;

import android.view.MotionEvent;
import androidx.lifecycle.Lifecycle;
import androidx.lifecycle.LifecycleObserver;
import androidx.lifecycle.OnLifecycleEvent;
import androidx.lifecycle.ProcessLifecycleOwner;

import tech.berty.gobridge.Logger;

public class MainActivity extends ReactActivity implements LifecycleObserver {
    private static final String TAG = "MainActivity";
    public enum AppState {
        Background,
        Foreground
    }

    private static AppState appState = AppState.Background;

    /**
     * Returns the name of the main component registered from JavaScript. This is
     * used to schedule rendering of the component.
     */
    @Override
    protected String getMainComponentName() {
        return "Berty";
    }

    public static AppState getAppState() {
        return MainActivity.appState;
    }

    @OnLifecycleEvent(Lifecycle.Event.ON_STOP)
    public void onAppBackgrounded() {
        //App in background
        MainActivity.appState = AppState.Background;
        Logger.d(TAG, "AppState: " + MainActivity.appState);
    }

    @OnLifecycleEvent(Lifecycle.Event.ON_START)
    public void onAppForegrounded() {
        //App in foreground
        MainActivity.appState = AppState.Foreground;
        Logger.d(TAG, "AppState: " + MainActivity.appState);
    }

    // needed by react-native-bootsplash
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // register for lifecycle events
        ProcessLifecycleOwner.get().getLifecycle().addObserver(this);
        RNBootSplash.init(R.drawable.bootsplash, MainActivity.this); // <- display the generated bootsplash.xml drawable
                                                                     // over our MainActivity
    }

    @Override
    protected void onStop() {
        super.onStop();
        MainActivity.appState = AppState.Background;
        Logger.d(TAG, "AppState: " + MainActivity.appState);
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
        MainActivity.appState = AppState.Background;
        Logger.d(TAG, "AppState: " + MainActivity.appState);
    }
}

  /**
   * Returns the instance of the {@link ReactActivityDelegate}. Here we use a util class {@link
   * DefaultReactActivityDelegate} which allows you to easily enable Fabric and Concurrent React
   * (aka React 18) with two boolean flags.
   */
  @Override
  protected ReactActivityDelegate createReactActivityDelegate() {
    return new DefaultReactActivityDelegate(
        this,
        getMainComponentName(),
        // If you opted-in for the New Architecture, we enable the Fabric Renderer.
        DefaultNewArchitectureEntryPoint.getFabricEnabled(), // fabricEnabled
        // If you opted-in for the New Architecture, we enable Concurrent React (i.e. React 18).
        DefaultNewArchitectureEntryPoint.getConcurrentReactEnabled() // concurrentRootEnabled
        );
  }
