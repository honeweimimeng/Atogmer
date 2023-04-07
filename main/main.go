package main

import "C"

func main() {
	//handle, err := windows.CreateIoCompletionPort(windows.InvalidHandle, 0, 0, 0)
	//if err != nil {
	//	panic(err.Error())
	//}
	//socket, err := windows.WSASocket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP, nil, 0, windows.WSA_FLAG_OVERLAPPED|windows.WSA_FLAG_NO_HANDLE_INHERIT)
	//if err != nil {
	//	panic(err.Error() + "===>")
	//}
	//handle, err = windows.CreateIoCompletionPort(socket, handle, uintptr(socket), 0)
	//if err != nil {
	//	panic(err.Error() + "===>")
	//}
	//err = windows.Bind(socket, &windows.SockaddrInet4{Port: 8080})
	//if err != nil {
	//	panic(err.Error() + "===>")
	//}
	//err = windows.Listen(socket, windows.SOMAXCONN)
	//if err != nil {
	//	panic(err.Error() + "===>")
	//}
	//for {
	//	newSocket, err := windows.WSASocket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP,
	//		nil, 0, windows.WSA_FLAG_OVERLAPPED|windows.WSA_FLAG_NO_HANDLE_INHERIT)
	//	if err != nil {
	//		panic(err.Error() + "===>")
	//	}
	//	windows.AcceptEx(socket, newSocket)
	//	if err != nil {
	//		panic(err.Error() + "===>")
	//	}
	//}
}
