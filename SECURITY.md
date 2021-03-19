# Security Policy

## Reporting a Vulnerability

The standard team and community take all security bugs in standard seriously. Thank you for improving the security of standard. We appreciate your efforts and responsible disclosure and will make every effort to acknowledge your contributions.

Report security bugs by emailing the lead maintainer at ernis.agilia@gmail.com

The lead maintainer will acknowledge your email within 48 hours, and will send a more detailed response within 48 hours indicating the next steps in handling your report. After the initial reply to your report, the security team will endeavor to keep you informed of the progress towards a fix and full announcement, and may ask for additional information or guidance.

Report security bugs in third-party modules to the person or team maintaining the module.

## Assurance Case Summary
Our overall security approach is called defense-in-breadth, that is, we consider security (including security countermeasures) in all our relevant software life cycle processes (including requirements, design, implementation, and verification). In each software life cycle process we identify the specific issues that most need to be addressed, and then address them.

There are other ways to organize assurance cases, and we have taken steps to ensure that issues that would covered by them are indeed covered. An alternate way to view security issues is to discuss "process, product, and people"; we evaluate the product in the verification process, and the people in the human resources process. It is important to secure the enabling environments, including the development environments and test environment; it may not be obvious, but that is covered by the infrastructure management process. At the end we cover certifications and controls, which also help us reduce the risk of failing to identify something important.

The following sections are organized following the assurance case figures:

We begin with the overall security requirements. This includes not just the high-level requirements in terms of confidentiality, integrity, and availability, but also access control in terms of identification, authentication (login), and authorization. Authentication is a cross-cutting and critical supporting security mechanism, so it's easier to describe it all in one place.
This is followed in the software life cycle processes, focusing on the software lifecycle technical processes: design, implementation, integration and verification, transition (deployment) and operations, and maintenance. We omit requirements, since that was covered earlier. This is a merger of the second and third assurance case figures (implementation is shown in a separate figure because there is so much to it, but in the text we merge the contents of these two figures).
We then discuss security implemented by other life cycle processes, broken into the main 12207 headings: agreement processes, organizational project-enabling processes, and technical management processes. Note that the organizational project-enabling processes include infrastructure management (where we discuss security of the development and test environment) and human resource management (where we discuss the knowledge of the key people involved in development).
We close with a discussion of certifications and controls. Certification processes can help us find something we missed, as well as provide confidence that we haven't missed anything important). Note that the project receives its own badge (the CII best practices badge), which provides additional evidence that it applies best practices that can lead to more secure software. Similarly, selecting IA controls can help us review important issues to ensure that the system will be adequately secure in its intended environment (including any compensating controls added to its environment). We controls in the context of the Center for Internet Security (CIS) Controls (aka critical controls).

https://github.com/AgiliaErnis/restaurateur/issues/5#issuecomment-802860466
![image](https://user-images.githubusercontent.com/40677903/111792508-9df55280-88c4-11eb-931a-9ed45d87e77b.png)
![image](https://user-images.githubusercontent.com/40677903/111792540-a51c6080-88c4-11eb-951b-9a2feb52a82f.png)
![image](https://user-images.githubusercontent.com/40677903/111792581-ad749b80-88c4-11eb-97bc-c5783f9664e6.png)
![image](https://user-images.githubusercontent.com/40677903/111792653-bcf3e480-88c4-11eb-8654-49a8b67e7169.png)
