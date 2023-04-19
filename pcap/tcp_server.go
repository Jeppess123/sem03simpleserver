
dekryptertMelding := mycrypt.Krypter([]rune(string(buf[:n]))), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
log.Println("Dekrypter melding: ", string(dekryptertMelding))
switch msg := string(dekrypterMelding) { ...
