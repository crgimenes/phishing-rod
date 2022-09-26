#include <M5Stack.h>


void setup() {
  M5.begin();
  M5.Power.begin();
  M5.Lcd.setTextSize(2);
  M5.Lcd.print("starting...\n");
}

void loop() {
  M5.update();
  if (M5.BtnA.pressedFor(1000, 200)) { // If A is pressed for 1 second power off
    M5.Power.powerOFF();
  }
}
