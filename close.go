package main



func (me *ChunkyFile) Close() error {

	//DEBUG
	me.log.Printf("[Close] BEGIN")	


	// Tell the go routine that we want to shutdown.
		done := make(chan struct{})

		me.shutdownWriter <- done


	// Wait for go routine to tell us it is done.

		//DEBUG
		me.log.Printf("[Close] Waiting for writer to shutdown.")	

		<-done


	//DEBUG
	me.log.Printf("[Close] END")	

	// Return.
		return nil
}
