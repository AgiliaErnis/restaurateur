# Security Policy

## Reporting a Vulnerability

The standard team and community take all security bugs in standard seriously. Thank you for improving the security of standard. We appreciate your efforts and responsible disclosure and will make every effort to acknowledge your contributions.

Report security bugs by emailing the lead maintainer at ernis.agilia@gmail.com

The lead maintainer will acknowledge your email within 48 hours and will send a more detailed response within 48 hours indicating the next steps in handling your report. After the initial reply to your report, the security team will endeavour to keep you informed of the progress towards a fix and full announcement and may ask for additional information or guidance.

Report security bugs in third-party modules to the person or team maintaining the module.


## Assurance Case Summary
Our overall security approach is called defence-in-breadth, that is, we consider security (including security countermeasures) in all our relevant software life cycle processes (including requirements, design, implementation, and verification). In each software life cycle process, we identify the specific issues that most need to be addressed, and then address them.

There are other ways to organize assurance cases, and we have taken steps to ensure that issues that would covered by them are indeed covered. An alternate way to view security issues is to discuss "process, product, and people"; we evaluate the product in the verification process, and the people in the human resources process. It is important to secure the enabling environments, including the development environments and test environment; it may not be obvious, but that is covered by the infrastructure management process. At the end we cover certifications and controls, which also help us reduce the risk of failing to identify something important.

The following sections are organized following the assurance case figures:

We begin with the overall security requirements. This includes not just the high-level requirements in terms of confidentiality, integrity, and availability, but also access control in terms of identification, authentication (login), and authorization. Authentication is a cross-cutting and critical supporting security mechanism, so it's easier to describe it all in one place.

This is followed in the software life cycle processes, focusing on the software lifecycle technical processes: design, implementation, integration and verification, transition (deployment) and operations, and maintenance. We omit requirements since that was covered earlier. This is a merger of the second and third assurance case figures (implementation is shown in a separate figure because there is so much to it, but in the text, we merge the contents of these two figures).

We then discuss security implemented by other life cycle processes, broken into the main 12207 headings: agreement processes, organizational project-enabling processes, and technical management processes. Note that the organizational project-enabling processes include infrastructure management (where we discuss security of the development and test environment) and human resource management (where we discuss the knowledge of the key people involved in development).

![image](https://user-images.githubusercontent.com/40677903/111792508-9df55280-88c4-11eb-931a-9ed45d87e77b.png)
![image](https://user-images.githubusercontent.com/40677903/111792540-a51c6080-88c4-11eb-951b-9a2feb52a82f.png)
![image](https://user-images.githubusercontent.com/40677903/111792581-ad749b80-88c4-11eb-97bc-c5783f9664e6.png)
![image](https://user-images.githubusercontent.com/40677903/111792653-bcf3e480-88c4-11eb-8654-49a8b67e7169.png)

## Confidentiality
## User Privacy
Part of our privacy requirement is that we "don't expose user activities to unrelated sites (including social media sites) without that user's consent"; here is how we do that.

We must first define what we mean by an unrelated site. A "related" site is a site that we are directly using to provide our service, in particular our cloud provider (Heroku which runs on Amazon's EC2 cloud-computing platform), authorization provider (Google), service provider (GitHub).

The steps we take steps to prevent unrelated sites from learning about our users' activities (and thus maintaining user privacy):

We directly serve all our own assets ourselves, including JavaScript, images, and fonts. In particular, we do not have any embedded automatically downloaded references (transclusions) in our web pages to external JavaScript or fonts. Since we serve these assets ourselves, and not via external third parties, external sites never receive any request from a user when they view our pages. As a result, user privacy is maintained: what a user views on our site is never revealed by our actions to unrelated sites. This also aids security; even if an attacker subverts some other site's JavaScript or font, that will not directly affect us because we do not embed references some other site's JavaScript or font in our web pages. Many sites don't do this and should probably consider it. This policy is enforced by our CSP (Content Security Policy).

We do not serve ads and we plan to have no ads in the future. That said, if we ever did serve ads, we expect that we would also serve them from our site, just like any other asset, to ensure that third parties did not receive unauthorized information.

We do not use any web analytics service that uses tracking codes or external assets. We log and store logs using only services we control or have a direct partnership with.

The email we send is privacy-respecting. The email contents we send do not have image links (which might expose when an email is read). In some cases, we have hyperlinks (e.g., to activate a local account), but those links go directly back to our site for that given purpose, and do not reveal information to anyone else. 

### User passwords
User passwords for local accounts are only stored on the server as iterated per-user salted hashes, and thus cannot be retrieved in an unencrypted form.

## Integrity
As noted above, HTTPS is used to protect the integrity of all communications between users and the application, as well as to authenticate the server to the user.

## Availability
As with any publicly accessible website, we cannot prevent an attacker with significant resources from temporarily overwhelming the system through a distributed denial-of-service (DDos) attack. So instead, we focus on various kinds of resilience against DDoS attacks and use other measures (such as backups in the future) to maximize availability. Thus, even if the system is taken down temporarily, we expect to be able to reconstitute it (including its data).

## Access Control
### Identification
Normal users must first identify themselves in one of two ways: (1) as a Google user with their Google account, or (2) as a custom "local" user with their email address.

The Restaurateur application runs on a deployment platform (Heroku), which has its own login mechanisms. Only those few administrators with deployment platform access have authorization to log in there, and those are protected by the deployment platform supplier (and thus we do not consider them further here). The login credentials in these cases are protected.

### Authentication
This system implements two kinds of users: local and remote. Local users log in using a password, but user passwords are only stored on the server as iterated salted hashes. Remote users use a remote system (we currently only support GitHub) using the widely used OAUTH protocol. 

A local user login will POST that information to /login, which is routed to session#create along with parameters such as session[email] and session[password]. If the encrypted hash of the password matches the stored hash, the user is accepted. If password doesn't match, the login is rejected. This is verified with these tests:
- Can login and edit using custom account
- Cannot login with local username and wrong password
- Cannot login with local username and blank password

A remote user login (pushing the "log in with Google" button) will invoke GET "/auth/google". The application then begins an omniauth login, by redirecting the user to Google login page. When the Google login completes, then per the omniauth spec there's a redirect back to our site.

### Authorization
Users who have not authenticated themselves can only perform actions allowed to anyone in the public (e.g., view the home page, view the list of restaurants, and view the information about each restaurant). Once users are authenticated, they are authorized to perform certain additional actions, like viewing history and suggestions.

## Assets
Our assets are:

- User passwords, especially for confidentiality. Unencrypted user passwords are the most critical to protect. As noted above, we protect these with encryption; we never store user passwords in an unencrypted or recoverable form.
- User email addresses, especially for confidentiality.
- Project data, primarily for integrity and availability. We back these up to support availability.

## Threat Agents
We have few insiders, and they are fully trusted to not perform intentionally hostile actions.

Thus, the threat agents we're primarily concerned about are outsiders, and the most concerning ones fit in one of these categories:

- People who enjoy taking over systems (without monetary benefit)
- Criminal organizations who want to take emails and/or passwords as a way to take over others' accounts (to break confidentiality). Note that our one-way iterated salted hashes counter easy access to passwords, so the most sensitive data is more difficult to obtain.
- Criminal organizations who want to destroy all our data and hold it for ransom (i.e., "ransomware" organizations). Note that our backups help counter this.
- Criminal organizations may try to DDoS us for money, but there's no strong reason for us to pay the extortion fee. We expect that people will be willing to come back to the site later if it's down, and we have scalability countermeasures to reduce their effectiveness. If the attack is ongoing, several of the services we use would have a financial incentive to help us counter the attacks. This makes the attacks themselves less likely (since there would be no financial benefit to them).

Like many commercial sites, we do not have the (substantial) resources necessary to counter a state actor who decided to directly attack our site. However, there's no reason a state actor would directly attack the site (we don't store anything that valuable), so while many are very capable, we do not expect them to be a threat to this site.

## DBMS
There is no direct access for normal users to the DBMS; in production, access requires special Heroku keys.

The DBMS does not know which user the Restaurateur is operating on behalf of and does not have separate privileges. However, the Restaurateur uses Active Record and prepared statements, making it unlikely that an attacker can use SQL injections to insert malicious queries.

- Spoofing identity. N/A, the database doesn't track identities.
- Tampering with data. The Restaurateur is trusted to make correct requests.
- Repudiation. N/A.
- Information disclosure. The Restaurateur is trusted to make correct requests.
- Denial of service. See earlier comments on DoS.
- Elevation of privilege. N/A, the DBMS doesn't separate privileges.

## Admin CLI
There is a command line interface (CLI) for admins. This is the Heroku CLI. Admins must use their unique credentials to log in.
- Spoofing identity. Every admin has a unique credential.
- Tampering with data. The communication channel is encrypted.
- Repudiation. Admins have unique credentials.
- Information disclosure. The channel is encrypted in motion.
- Denial of service. Heroku has a financial incentive to keep this available and takes steps to do so.
- Elevation of privilege. N/A; anyone allowed to use this is privileged.

## Secure design principles
Applying various secure design principles helps us avoid security problems in the first place. The most widely used list of security design principles, and one we build on, is the list developed by Saltzer and Schroeder.

Here are several secure design principles and how we follow them, including all 8 principles from Saltzer and Schroeder:
- Separation of privilege (multi-factor authentication, such as requiring both a password and a hardware token, is stronger than single-factor authentication): We don't use multi-factor authentication because the risks from compromise are smaller compared to many other systems (it's almost entirely public data, and failures generally can be recovered through backups).
- Least privilege (processes should operate with the least privilege necessary).
- Least common mechanism (the design should minimize the mechanisms common to more than one user and depended on by all users, e.g., directories for temporary files): No shared temporary directory is used. Each time a new request is made, new objects are instantiated; this makes the program generally thread safe as well as minimizing mechanisms common to more than one user. The database is shared, but each table row has access control implemented which limits sharing to those authorized to share.
- Psychological acceptability (the human interface must be designed for ease of use, designing for "least astonishment" can help): The application presents a simple login and "fill in the form" interface, so it should be acceptable.
- Limited attack surface (the attack surface, the set of the different points where an attacker can try to enter or extract data, should be limited): The application has a limited attack surface. 
- Input validation with allow lists (inputs should typically be checked to determine if they are valid before they are accepted; this validation should use allow lists (which only accept known-good values), not deny lists (which attempt to list known-bad values)): In data provided directly to the web application, input validation is done with allow lists through controllers and models.

## Memory-safe Languages
Golang and JavaScript are both memory-safe languages, so the vulnerabilities of memory-unsafe languages (such as C and C++) cannot occur in the custom code. This also applies to most of the code in the directly depended on libraries.

## Security in Implementation
Most implementation vulnerabilities are due to common types of implementation errors or common misconfigurations, so countering them greatly reduces security risks.

To reduce the risk of security vulnerabilities in implementation we have focused on countering the OWASP Top 10, both the OWASP Top 10 (2013) and OWASP Top 10 (2017).
.