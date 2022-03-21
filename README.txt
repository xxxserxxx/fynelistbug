1. Build the project
   ```sh
   fyne package -os android
   ```
2. Run your emulator
   ```sh
   emulator -avd pixel5
   ```
3. Install the project
   ```sh
   adb -s emulator-5554 install Fyne_List_Bug.apk
   ```
4. Log the emulator
   ```sh
   adb -s emulator-5554 logcat
   ```
   optionally for the test
   ```sh
   adb -s emulator-5554 logcat | grep -i fyne
   ```
5. Start the process (either through the UI, or with this command)
   ```sh
   adb -s emulator-5554 shell am start -n net.ser1.fynelistbug/org.golang.app.GoNativeActivity
   ```
6. Check the directory via the shell:
   ```sh
   adb -s emulator-5554 shell ls -a /data/user/0/net.ser1.fynelistbug/files/fyne/fynelistbug/subdir
   ```

The call to `List()` should have returned an empty array; instead, it returns an array with a nil. The message in the UI will show 1 nil, 0 non-nil.
