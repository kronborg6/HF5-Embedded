void setup() {
/*   char temp "20"
  char hum "20" */
  // put your setup code here, to run once:
Serial.begin(9600);
  
  // Serial.println("test");

  // api("POST", "/data", "\r\n{\"temp\":10.2,\r\n\"hum\":24.4,\r\n\"local_id\":2\r\n  \r\n}\r\n");
  String test = api("GET", "/startup/2", "");
  Serial.println("Her vil jeg se data");
  Serial.println(test);
}

void loop() {
  // put your main code here, to run repeatedly:

}
