# GRCHive v2

This is the monorepository containing all the code to support our products:

* GRCHive Base
* GRCHive Workflow
    * GRCHive Scoping
    * GRCHive PBC
* GRCHive Data Repository
* GRCHive Audit Automation
* GRCHive Audit ETL
* GRCHive Continuous Auditing & Monitoring

Note that the entire product should be built in a manner that can be used by external and internal audit teams.
Our suite of products is meant to compete with the likes of:

* HighBond
* Audit Board
* Logic Gate

## GRCHive Base

The Base product contains common features that will be required by all other GRCHive products.
This includes:

* Notifications
* Role-based Access Controls
* Encryption/Decryption
* Scheduled and Repeatable Tasks
* Comments
* Audit Trail

## GRCHive Workflow

Our Workflow product should provide the tools necessary to complete an audit from start to finish with the assumption that the general audit workflow is:

1. Scoping: Determine what "objects" are in scope for the audit.
1. PBC: For the given "object", request relevant data which could be documentation or samples.
1. Test: Ensure controls are designed and working properly using the data collected in the PBC process.
1. Review: Allow higher-level employees to review and sign-off on the work to ensure that the work done is up to standards.
1. Remediation: If a test or review fails, provide steps to fix the problem and redo the audit of this particular "object."

Furthermore, the Workflow product should provide the tools to track the overall progress of the audit as well as progress within each particular step (if relevant).
All the products under the Workflow product will live within the concept of an "engagement" which is just how all the information gathered and work performed during a specific audit engagement is tracked.
An audit "engagement" can be duplicated to roll-forward.

### GRCHive Scoping

The Scoping product should allow auditors to collect information about any object that may be in scope for their audit.
Currently, our planned list of supported objects are:

* Risks
* Controls
* Vendors
* General Ledger Accounts
* Databases
* Systems
* Servers
* (TBD CMMC) Network Devices (e.g. VPN servers, routers, firewalls, wireless access points, etc.)
* (TBD CMMC) Computer Inventory (e.g. Desktops, Laptops, etc.)

We will provide functionality to allow users to create and edit these objects on GRCHive with specialized tools where necessary.
For example, we will allow the creation of Process Flows via an interactive graph editor similar to Visio.

Furthermore, we will allow users to seamlessly view and create relationships between any of our supported objects.
These relationships will either be a direct relationship (e.g. directly specifying a risk and control are related) or a graph relationship.
Graph relationships will exist in the form of:

* Process Flows
* (TBD CMMC) Network Map

These graph relationships will also be considered potential in-scope object(s).

**Competing Products**: Visio, Excel

**Competitive Advantage**: 

* Easy to keep relationships in sync (i.e. keeping process flows in sync with risks and controls).
* Centralized source of truth for all objects in scope.

### GRCHive PBC

The PBC product builds upon the scoping product and allows auditors to create PBC requests for **any** of the supported objects.
Alternatively, the PBC request can be completely independent of any of the in scope objects.
We will currently plan to support two types of PBC requests:

* File Requests
* IT-Integrated Requests

File requests will be PBC requests where the auditee is expected to upload one or more files.
IT-integrated requests will be PBC requests where the auditor directly requests information from one of the IT-related in-scope objects (e.g. running a SQL query on a database) using the Audit ETL product.
Note that all PBC request results will be stored as files in the Data Repository product.
IT-integrated requests will be stored as "smart" files which will be specialized files that contain not only the data returned from the IT-object but also what command was used to obtain that data.

All PBC requests will go through the following stages:

1. Open: When the PBC request is created by an auditor.
1. In Progress: When the PBC request is being worked on by the auditee (i.e. starting to upload files or allowing the IT-integrated request to run).
1. Pending Approval: When the auditee determines that they have completed the PBC request.
1. Feedback: If the auditor determines the documents uploaded are lacking, they will provide feedback and ask the auditee to re-upload files.
1. Approved: If the auditor determines the documents uploaded are sufficient, they will mark the PBC as approved.

**Competing Products**: Email + Sharefile/Dropbox/Google Drive/etc., Suralink, AuditFile

**Competitive Advantage**:

* Centralized location to keep track of PBCs means no more lost/duplicated PBC requests.
* Automated notifications means that auditors no longer need to manually follow up on requests.
* Specialized PBCs allow auditors to specifically request what they want (e.g. SQL query) instead of relying on auditees to accurately read, interpret, and respond to the request.
* Integration with GRCHive Staging for better organization.
* Progress dashboard (not an advantage over Suralink, AuditFile).

## GRCHive Data Repository

The primary purpose of the Data Repository is to act as a versioned file system for all the files that will be uploaded throughout the course of multiple audits.
Each file will be versioned so that if a user decides to upload a new version of the file, both the old file and the newly uploaded file will still be accessible.
We will provide PDF previews on the website.

Files will be primarily organized using folders as similar to common desktop operating systems.
For every organization, there will be a root organization folder which keeps track of every document ever uploaded for this organization.
For every audit engagement, there will be an engagement folder where all the documents for the engagement will live.
These folders will live under the organization folder.
Furthermore, for every in-scope object, there will be a folder where all the documents for that object will live.
These folders will live under the engagement folder.
Aside from these default folders associated with concepts (engagements, in-scope objects) in GRCHive, users are free to edit/delete/create folders; however, unless the file lives in the in-scope object folder (or a sub-folder of), the file will not be possible when visiting the overview page of the in-scope object.

Note that PBC requests will be directly tied to **folders** and **folders** will be directly tied to in-scope objects.
This easily gives PBC requests the flexibility to be used for any in-scope object and allows users to organize files easily as PBC requests can then put files into the desired subfolder for organizational purposes.

**Competing Products**: Dropbox, Box, Google Drive, FTP server

**Competitive Advantage**:

* Integration with PBC and Scoping allows for in-context viewing of documents on the site.
* Integration with Base allows for comments directly in the context of the document.

## GRCHive Audit ETL

Audit ETL is the suite of connectors that will be responsible pulling data from client databases, SaaS apps, servers etc. directly to GRCHive.
Audit ETL will also include a library of pre-built queries/commands for pulling relevant information from IT systems to minimize the amount of technical knowledge needed by auditors (e.g. user listings, configuration data).
PBC will utilize Audit ETL to extract data out from relevant IT systems contained in Scoping and store that information as a "smart" file in Data Repository.

**Competing Products**: FiveTrans, Stitch (Talend), Hevo, other ETL/ELT products

**Competitive Advantage**:

* Integration with PBC allows us to retain audit-relevant information along with the data.
* Support for commonly audited platforms (e.g. SAP, Intuit Quickbooks) that may or may not be supported by commercial ETL/ELT products.

## GRCHive Audit Automation

Customizable scripts to automatically perform tests.

**Competing Products**:  ElectroNeek, Linx, Microsoft Power Automate, Custom scripts

**Competitive Advantage**:

* Tailored for audit-specific applications.

## GRCHive Continuous Auditing & Monitoring

TBD.

**Competing Products**: None.

**Competitive Advantage**: N/A.
