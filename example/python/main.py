import socket
import time


def send_tcp_message(host: str, port: int):
    """Send a message to the OpenProxy proxy via TCP socket every second."""
    try:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client_socket:
            client_socket.connect((host, port))
            print(f"Connected to {host}:{port}")

            while True:
                # Send a message
                data_ownercode = "CXX"
                block_code = "40"
                vehicle_number = "2046"
                latitude = 52.38744 # WGS84 latitude
                longitude = 4.6369134  # WGS84 longitude
                accuracy = 4.0  # Accuracy in meters
                speed = 3.1  # Speed in m/s
                bearing = 130.3  # Bearing in degrees
                odometer = 13242  # Odometer in meters
                hdop = 4  # Horizontal dilution of precision
                number_of_received_satellites = 16
                timestamp = 1613652529565  # Milliseconds since epoch
                dooropeningstatus = 2  # Example value from .proto 0=NO_DATA, 1=CLOSED, 2=OPEN, 3=RELEASED
                stopbuttonstatus = 2  # Example value from .proto 0=NO_DATA, 1=NOT_ACTIVATED, 2=ACTIVATED
                number_of_vehicles_coupled = 1  # Example value from .proto
                driving_direction = ""  # Could be "nothing", "A", or "B"
                passage_stop_status = 2  # Example value from .proto 0=UNKNOWN, FALSE=1, TRUE=1

                # Combine the variables into a single CSV string
                message = ",".join(map(str, [
                    data_ownercode, block_code, vehicle_number, latitude, longitude,
                    accuracy, speed, bearing, odometer, hdop, number_of_received_satellites,
                    timestamp, dooropeningstatus, stopbuttonstatus, number_of_vehicles_coupled,
                    driving_direction, passage_stop_status
                ])) + "\n"

                client_socket.sendall(message.encode('utf-8'))
                print("Sent: " + message)

                # Wait for 1 second
                time.sleep(1)

    except ConnectionRefusedError:
        print("Connection refused. Ensure the server is running and accessible.")
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    # Define the server details
    HOST = "127.0.0.1"  # Replace with the target server's IP address
    PORT = 9001         # Replace with the target server's port number

    send_tcp_message(HOST, PORT)