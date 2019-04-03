package main




var shouldQuit=make(chan struct{})
func main(){
{
//loop
}
//...out of the loop
select {
case <-c.shouldQuit:
cleanUp()
return
default:
}
//...
}