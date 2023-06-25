#include <ESP8266WiFi.h>
#include <WiFiUdp.h>
#include <TFT_eSPI.h>
#include <SPI.h>

/* HIDDEN */
const char *const ssid = "hidden";
const char *const pass = "hidden";
const IPAddress serverIP = {192, 168, 1, 80};

const int serverPort = 1050;
char orderReceived[2] = "0";

WiFiUDP UDP;
TFT_eSPI screen = TFT_eSPI();

void setup() {
  screen.init();
  screen.setRotation(3);
  screen.setTextSize(1);

  
  screen.fillScreen(TFT_BLACK);
  screen.setCursor(10, 10);
  screen.println("Connecting to WiFi");
  
  
  int status = WL_IDLE_STATUS;
  while (status != WL_CONNECTED) {
    status = WiFi.begin(ssid, pass);
    delay(10000);
  }
  UDP.begin(1050);

  screen.fillScreen(TFT_BLACK);
  screen.setCursor(10, 10);
  screen.println(WiFi.localIP());
  screen.setCursor(5, 40);
  screen.println("Waiting for assignment.");
  screen.setCursor(10, 30);

	/* Recover the info in case it was a small poweroff */
  UDP.beginPacket(serverIP, serverPort);
  UDP.write("2");
  UDP.endPacket();
  delay(5000);
}

void loop() {
  int packetSize = UDP.parsePacket();
  if (packetSize) {
    orderReceived[0] = '1';
    char *name = (char *)alloca(packetSize + 1);

    UDP.read(name, packetSize);
    name[packetSize] = '\0';

    char *order;
    for (int i = 0; i < packetSize; ++i) {
      if (name[i] == '\n') {
        order = name + i + 1;
        name[i] = '\0';
      }
    }

    screen.fillScreen(TFT_BLACK);
    screen.setCursor(10, 10);
    screen.println(name);
    screen.setCursor(5, 40);
    screen.println(order);
  } else {
    orderReceived[0] = '0';
  }

  UDP.beginPacket(serverIP, serverPort);
  UDP.write(orderReceived);
  UDP.endPacket();
  delay(5000);
}

