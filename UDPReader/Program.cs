using System;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using Newtonsoft.Json;
using Kaitai;

namespace UDPReader
{
    class Program
    {
        const int PORT = 4000;
        static Thread doWork;

        static void Main(string[] args)
        {   
            doWork = new Thread(new ThreadStart(startReading));

            doWork.Start();

            Console.Read();
        }

        static void OnReaderMessage(Byte[] bytes) {
            AhrsRecord ahrs = new AhrsRecord(new KaitaiStream(bytes));
            Console.WriteLine(ahrs.Record.ToString());        
        }

        static void startReading() {
            DgramReader reader = new DgramReader();
            reader.OnMessageReceived += OnReaderMessage;
            reader.Run(PORT);
        }
    }

    class DgramReader 
    {
        public delegate void OnMessageReceivedDelegate(Byte[] message);
        public OnMessageReceivedDelegate OnMessageReceived;

        public void Run(int port) {
            
            UdpClient receivingUdpClient = new UdpClient(port);

            IPEndPoint RemoteIpEndPoint = new IPEndPoint(IPAddress.Any, 0);
            
            try {
                while (true) {
                    Byte[] buffer = receivingUdpClient.Receive(ref RemoteIpEndPoint);

                    if (buffer[1] == 76) {
                        OnMessageReceived?.Invoke(buffer);
                    }
                }
            }
            catch ( Exception e ) {
                Console.WriteLine(e.ToString());
            }
        }
    }
}