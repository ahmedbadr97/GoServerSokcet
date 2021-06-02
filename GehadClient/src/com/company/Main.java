package com.company;

import java.io.IOException;
import java.io.PrintWriter;
import java.net.Socket;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        try {
            Scanner cliInput=new Scanner(System.in);
            Socket  socket=new Socket("127.0.0.1",5000);
            System.out.println("Please help us finding gehad \n where is the last location you have seen her at:");
            PrintWriter outputToSocket=new PrintWriter(socket.getOutputStream());
            outputToSocket.println(cliInput.nextLine());
            System.out.println("Location sent to Detective conan \nThank you for your collaboration \n");
            outputToSocket.close();
            socket.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
        // write your code here
    }
}
