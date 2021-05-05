# CYBER HYGIENE (BITAPPLE) V1.0

## CHANGE LOG

| TIME     | BY        | VERSION | DESCRIPTION     |
| -------- | --------- | ------- | --------------- |
| 2020-3-1 | zhoujiabo | V1.0    | Create Document |
|          |           |         |                 |



## OVERVIEW

### The Shared Responsibility Model

![](/src/amas/docs/source/_drafts/cyber_hygiene.assets/DeepinScreenshot_select-area_20201119103558-1605774418114.png)

----



### Administrative Accounts

```
Secure the use of every administrative account in respect of any 
operating system, database, application, security appliance or network device through preventive controls. These controls should prevent the unauthorised access to or use of such account.
```

#### AWS Root Accounts

The most privileged user account in an AWS account is the [root user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html). The root user is associated with the provided email address and password used to create the account. The root user account has access to every resource in the account—including the ability to close it. To align to the principle of [least privilege](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#grant-least-privilege), the root user account  ***not*** be used for everyday tasks. Instead, [AWS Identity and Access Management (IAM) roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)  created and scoped to particular roles and functions to do actual works. 我们遵循: [Security best practices in IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)

- Use  strong password (Both register Email & Root User Account)
- Removed any programmatic access keys from the root user account
- The root user password  store in 2 Hardware encrypted(AES256) USB stick with separate permissions for access.

- [Enable multi-factor authentication (MFA)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#id_root-user_manage_mfa), and use U2F hardware security keys. Using 2 separate hardware encrypted(AES256) USB stick store for the password and the MFA token, with separate permissions for access.

#### IAM User Accounts

We use IAM user do do normal workd and maintenance.

- Use groups to assign permissions to IAM users
- Grant least privilege
- We followed [AWS managed policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html#aws-managed-policies) 

![         Diagram of AWS managed policies       ](https://docs.aws.amazon.com/IAM/latest/UserGuide/images/policies-aws-managed-policies.diagram.png)



#### Database Accounts

- DBA Accounts
  - Use IAM database authentication
- Application Accounts
  - Use AWS KMS to entryped store & manage the password used by application
  - Rotating keys(Encryped Password) by for a period of time
  - Controlling KMS key access
  - Auditing KMS key usage and lifecycles
  - Grant least privilege
    - read permission
    - read/write permisson

#### Server Accounts

- Bastion Machine network(VPC) is isolated from Normal App Server.

- Use Bastion Machine as the entrance of the server(EC2)
- MFA(HOTP) is required to login basion machine
- Login with special work user to perform normal operation
- Server root privileged need approved and the privileged revoked in 8 hours automatically

###  Security Patches

- A System Owner or team must be identified for the overall security management of each system or device. 
- Those responsible for each system, device and application must monitor relevant sources of information which may alert them to a need to act in relation to new security vulnerabilities.
- Patches must be obtained from a known, trusted source.
- Patches must be tested and assessed before implementation in a production environment to ensure that there is no negative impact as a result.
- A backup of the production systems must be taken before applying any patch.
- An audit trail of all changes must be created and documented. The System
  Owner must verify that the patches have been installed successfully after
  production deployment.
- Production patches must be deployed regularly as per the SLA defined below.
- System owners outside of IT Services that manage the security of their own
  systems are required to use patches in accordance with this procedure.

- Any os / libraries / source code which to build service are managed by universal repository manager
- The integrity of patches must be verified through such means as comparisons of cryptographic hashes to ensure the patch obtained is the correct, unaltered patch.
- Only the artifacts in  universal repository manager can be use for building our service
- Tracing the artifacts where to be used
- Scaned artifacts in universal repository  for a period of time base on public vulnerability database (e.g. CVE, NVD)
- OS security patch
  - Use AWS Systems Manager Patch Manager to apply patch to the OS
- Software security patch
  - Remove the risk artifacts from universal repository and rebuild service



SLA with Priority:

![](/src/amas/docs/source/_drafts/cyber_hygiene.assets/DeepinScreenshot_select-area_20201119152351.png)

###  Security Standards
We build services on AWS cloud. The following is a partial list of assurance programs with which AWS complies:

- SOC 1/ISAE 3402, SOC 2, SOC 3
- FISMA, DIACAP, and FedRAMP
- PCI DSS Level 1
- ISO 9001, ISO 27001, ISO 27017, ISO 27018 

#### ISO 27001 

ISO 27001 is a security management standard that specifies
security management best practices and comprehensive security controls
following the ISO 27002 best practice guidance. The basis of this certification is
the development and implementation of a rigorous security program, which
includes the development and implementation of an Information Security
Management System which defines how AWS perpetually manages security in a
holistic, comprehensive manner. For more information, or to download the AWS
ISO 27001 certification, see https://aws.amazon.com/compliance/iso-27001-faqs/

#### ISO 27017 
ISO 27017 provides guidance on the information security aspects of
cloud computing, recommending the implementation of cloud-specific information
security controls that supplement the guidance of the ISO 27002 and ISO 27001
standards. This code of practice provides additional information security controls
and implementation guidance specific to cloud service providers. For more
information, or to download the AWS ISO 27017 certification, see
https://aws.amazon.com/compliance/iso-27017-faqs/

#### ISO 27018

ISO 27018 is a code of practice that focuses on protection of
personal data in the cloud. It is based on ISO information security standard
27002 and provides implementation guidance on ISO 27002 controls applicable
to public cloud Personally Identifiable Information (PII). It also provides a set of
additional controls and associated guidance intended to address public cloud PII
protection requirements, which is not addressed by the existing ISO 27002
control set. For more information, or to download the AWS ISO 27018
certification, see https://aws.amazon.com/compliance/iso-27018-faqs/

#### ISO 9001

ISO 9001 outlines a process-oriented approach to documenting and
reviewing the structure, responsibilities, and procedures required to achieve
effective quality management within an organization. The key to the ongoing
certification under this standard is establishing, maintaining and improving the
organizational structure, responsibilities, procedures, processes, and resources
in a manner in which AWS products and services consistently satisfy ISO 9001
quality requirements. For more information, or to download the AWS ISO 9001
certification, see https://aws.amazon.com/compliance/iso-9001-faqs/

#### MTCS Level 3 

Multi-Tier Cloud Security (MTCS) is an operational Singapore
security management Standard (SPRING SS 584:2013), based on ISO 27001/02
Information Security Management System (ISMS) standards. The key to the
ongoing three-year certification under this standard is the effective management
of a rigorous security program and annual monitoring by an MTCS Certifying
Body (CB). The Information Security Management System (ISMS) required under
this standard defines how AWS perpetually manages security in a holistic,
comprehensive way. For more information, see
https://aws.amazon.com/compliance/aws-multitiered-cloud-security-standard-
certification/

#### PCI DSS Level 1

The Payment Card Industry Data Security Standard (also
known as PCI DSS) is a proprietary information security standard administered
by the PCI Security Standards Council. PCI DSS applies to all entities that store,
process, or transmit cardholder data (CHD) and/or sensitive authentication data
(SAD) including merchants, processors, acquirers, issuers, and service
providers. The PCI DSS is mandated by the card brands and administered by the
Payment Card Industry Security Standards Council. For more information, or to
request the PCI DSS Attestation of Compliance and Responsibility Summary,
see https://aws.amazon.com/compliance/pci-dss-level-1-faqs/

#### SOC
AWS Service Organization Control (SOC) Reports are independent, third-party examination reports that demonstrate how AWS achieves key compliance
controls and objectives. The purpose of these reports is to help customers and
their auditors understand the AWS controls established to support operations
and compliance. For more information, see
https://aws.amazon.com/compliance/soc-faqs/
There are three types of AWS SOC Reports:

- SOC 1 – Provides information about the AWS control environment that might
  be relevant to a customer’s internal controls over financial reporting as well as
  information for assessment and opinion of the effectiveness of internal
  controls over financial reporting (ICOFR).
- SOC 2 – Provides customers and their service users that have a business
  need with an independent assessment of the AWS control environment that is
  relevant to system security, availability, and confidentiality.
- SOC 3 – Provides customers and their service users that have a business
  need with an independent assessment of the AWS control environment that is
  relevant to system security, availability, and confidentiality, without disclosing
  AWS internal information

### Network Perimeter Defence
Implement controls at its network perimeter to restrict all unauthorised network traffic. 


###  Malware Protection
Implement malware protection measures on every system to mitigate the risk of malware infection, where available and can be implemented. 

### Multi-factor Authentication
Strengthen user authentication through implementation of multifactor authentication for all administrative accounts in respect of any operating system, database, application, security appliance or network device that is a critical system, and all accounts on any system used to access customer information through the internet.



## Appendix

- AWS User Guide to Financial Services Regulations & Guidelines in Singapore
- [AWS Artifact](https://aws.amazon.com/artifact/?nc1=h_ls)
- [Risk and Compliance Whitepaper](https://d0.awsstatic.com/whitepapers/compliance/AWS_Risk_and_Compliance_Whitepaper.pdf)
- [Overview of Security Process Whitepaper](https://d0.awsstatic.com/whitepapers/aws-security-whitepaper.pdf)
- [SOC audit reports](https://aws.amazon.com/compliance/soc-faqs/)
- [AWS Well-Architected Framework](https://aws.amazon.com/architecture/well-architected/)
- [AWS Config managed rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) 
- https://aws.amazon.com/compliance/
- https://assets.publishing.service.gov.uk/government/uploads/system/uploads/attachment_data/file/882770/dwp-ss015-security-standard-malware-protection-v1.1.pdf