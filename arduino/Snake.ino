void setup(){
  Serial.begin(115200);
}

void loop(){
  Serial.write(1);
  delay(1000);
}
