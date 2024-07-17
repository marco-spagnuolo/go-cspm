# Go CSPM

**Go-CSPM** (Go Cloud Security Posture Management) is a lightweight, Go-based tool designed to help organizations maintain secure cloud environments. It automates the process of checking cloud infrastructure against industry standards and best practices, providing insights and recommendations for improving security postures.

## Features

- **Automated Security Checks**: Run predefined security checks on your cloud infrastructure.
- **Customizable Rules**: Define custom security rules based on your organization's requirements.
- **Report Generation**: Generate comprehensive security reports with actionable recommendations.
- **Extensible**: Easily extend the tool with additional checks and functionalities.
- **Lightweight**: Minimal dependencies, easy to deploy and run.

## Installation

To install Go-CSPM, you need to have Go installed on your machine. Then run:
```bash
go get github.com/marco-spagnuolo/go-cspm
```


## Generating Reports

To generate a security report, use:

```bash
go-cspm -report
```

## Custom Rules

You can define custom rules by creating a rules file in YAML format. Here is an example:

```yaml
- id: custom-rule-1
  description: "Ensure that all S3 buckets are encrypted"
  check: "s3:CheckBucketEncryption"

- id: custom-rule-2
  description: "Ensure there are no unused IAM access keys"
  check: "iam:CheckUnusedKeys"

- id: custom-rule-3
  description: "Ensure EC2 instances are of approved types"
  check: "ec2:CheckInstanceTypes"

- id: custom-rule-4
  description: "Ensure RDS instances have backup retention policy enabled"
  check: "rds:CheckBackupRetention"
```
## To run checks with your custom rules, use:

```bash 
go-cspm -check -rules /path/to/custom/rules.yaml
```

Contributing
Contributions are welcome! If you would like to contribute to Go-CSPM, please follow these steps:

    Fork the repository.
    Create a new branch (git checkout -b feature-branch).
    Make your changes.
    Commit your changes (git commit -am 'Add new feature').
    Push to the branch (git push origin feature-branch).
    Create a new Pull Request.

License
Go-CSPM is licensed under the MIT License. See the LICENSE file for more details.
Contact
For any questions or inquiries, please open an issue on the GitHub repository.