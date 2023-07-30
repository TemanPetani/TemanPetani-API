# TemanPetani Project

TemanPetani Project is a web application that aims to provide farmers with easy access to agricultural resources. With features such as user management, product listings, sales and purchases, negotiation options, farming activity tracking, and notifications, it facilitates efficient farming operations and resource management. Admins can add agricultural products for sale, while farmers can showcase their produce and engage in negotiation for fair pricing. This application is built using Golang with Echo, Gorm, GCP, Docker, and Midtrans as payment gateway.

## Key Features

1. **User Management:** The application provides user registration and authentication functionalities for farmers and admins.

2. **Product Listings:** Admins can add agricultural products to the platform for sale, providing a wide range of options for farmers and buyers.

3. **Sales and Purchases:** Farmers can showcase their produce on the platform, while buyers can make purchases directly from the website.

4. **Negotiation Options:** The platform enables farmers and buyers to engage in negotiations, ensuring fair pricing and successful transactions.

5. **Farming Activity Tracking:** Farmers can track their farming activities and progress through the application, allowing for efficient monitoring and planning.

6. **Notifications:** The platform sends timely notifications to users, keeping them informed about important updates, negotiations, and transactions.

## Technology Stack

- Golang
- Echo (Web framework)
- Gorm (ORM library)
- GCP (Google Cloud Platform)
- Docker (Containerization)
- Midtrans (Payment gateway)

## Getting Started

To get started with the TemanPetani Project, follow the steps below:

1. Clone the repository: `git clone https://github.com/TemanPetani/TemanPetani.git`

2. Install the required dependencies by running: `go mod download`

3. Set up the database and environment variables accordingly.

4. Build and run the application: `go run main.go`

## Contributing

Contributions to the TemanPetani Project are welcome! If you have any suggestions, bug reports, or feature requests, please feel free to open an issue or submit a pull request.

## License

TemanPetani Project is open-source and distributed under the [MIT License](LICENSE).

## Contact

For any inquiries or further information, you can reach us at contact@temanpetani.com.

Thank you for choosing TemanPetani Project! Happy farming! 🌾🚜


## Local.env

The `local.env` file contains environment variables required to configure your application when running in a local development environment. This file includes configuration values such as database settings, access tokens, and more.

- Configuration

To use the `local.env` file, follow these steps:

1. Create a file named `local.env` in your project directory.
2. Fill in the empty values with the appropriate configuration for your local environment:

```plaintext
DB_USERNAME=
DB_PASS=
DB_HOSTNAME=
DB_PORT=
DB_NAME=
JWT_SECRET_KEY=
