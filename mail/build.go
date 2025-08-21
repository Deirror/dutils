package mail

import "fmt"

// EmailBuild constructs HTML email bodies for verification and notification emails.
// It holds the base domain used for building verification links.
type EmailBuild struct {
	domain string
}

// NewEmailBuild creates a new EmailBuild with the given domain as the base URL.
// The domain should include scheme (e.g., https://example.com) and optionally a trailing slash.
func NewEmailBuild(domain string) *EmailBuild {
	return &EmailBuild{
		domain: domain,
	}
}

// VerificationEmail returns an HTML email body for verifying a user's email address after registration.
// 'callback' is the verification endpoint path appended to the domain.
// 'code' is the unique verification code appended as a query parameter.
func (b *EmailBuild) VerificationEmail(callback, code string) string {
	verifyLink := fmt.Sprintf("%s%s?code=%s", b.domain, callback, code)

	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Verify Your Email</title>
		</head>
		<body style="font-family: sans-serif; background: #f4f4f4; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background: white; padding: 20px; border-radius: 8px;">
				<h2 style="color: #333;">Verify your email address</h2>
				<p>Thanks for registering in our system! Please click the button below to verify your email address.</p>
				<p style="text-align: center; margin: 24px 0;">
					<a href="%s" style="background: #4f46e5; color: white; padding: 12px 24px; text-decoration: none; border-radius: 4px;">
						Verify Email
					</a>
				</p>
				<p>If the button doesn't work, paste this link in your browser:</p>
				<p><a href="%s">%s</a></p>
				<p style="font-size: 12px; color: #666;">This message was sent automatically. Please do not reply directly to this email.</p>
			</div>
		</body>
		</html>
	`, verifyLink, verifyLink, verifyLink)
}

// VerificationChangeEmail returns an HTML email body for verifying a user's new email address after an email change.
// 'callback' is the verification endpoint path appended to the domain.
// 'code' is the unique verification code appended as a query parameter.
func (b *EmailBuild) VerificationChangeEmail(callback, code string) string {
	verifyLink := fmt.Sprintf("%s%s?code=%s", b.domain, callback, code)

	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Verify Your Email</title>
		</head>
		<body style="font-family: sans-serif; background: #f4f4f4; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background: white; padding: 20px; border-radius: 8px;">
				<h2 style="color: #333;">Verify your email address</h2>
				<p>Thanks for using our system! Please click the button below to verify your new email address.</p>
				<p style="text-align: center; margin: 24px 0;">
					<a href="%s" style="background: #4f46e5; color: white; padding: 12px 24px; text-decoration: none; border-radius: 4px;">
						Verify Email
					</a>
				</p>
				<p>If the button doesn't work, paste this link in your browser:</p>
				<p><a href="%s">%s</a></p>
				<p style="font-size: 12px; color: #666;">This message was sent automatically. Please do not reply directly to this email.</p>
			</div>
		</body>
		</html>
	`, verifyLink, verifyLink, verifyLink)
}

// NotifyOldEmail returns an HTML email body notifying the user at their old email address
// that their account email was changed.
// 'sysName' is the name of the system or application.
// 'support' is the support email address to contact if the change was unauthorized.
func (b *EmailBuild) NotifyOldEmail(sysName, support string) string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Email Change Notification</title>
	</head>
	<body style="font-family: sans-serif; background: #f4f4f4; padding: 20px;">
		<div style="max-width: 600px; margin: auto; background: white; padding: 20px; border-radius: 8px;">
			<h2 style="color: #333;">Your email was changed</h2>
			<p>This is a security notice from <strong>%s</strong>. Your account’s email address has recently been changed.</p>
			<p>If you made this change, no further action is needed.</p>
			<p>If you <strong>did not</strong> make this change, please contact our support team immediately at 
			<a href="mailto:%s">%s</a>.</p>
			<hr style="margin: 24px 0; border: none; border-top: 1px solid #ddd;">
			<p style="font-size: 12px; color: #666;">This message was sent automatically. Please do not reply directly to this email.</p>
		</div>
	</body>
	</html>
`, sysName, support, support)
}

// VerificationDeleteAccount returns an HTML email body for verifying a user's account deletion request.
// 'sysName' is the system name, 'support' is the support email, 'callback' is the endpoint path, and 'code' is the unique verification code.
func (b *EmailBuild) VerificationDeleteAccount(sysName, support, callback, code string) string {
	verifyLink := fmt.Sprintf("%s%s?code=%s", b.domain, callback, code)

	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Confirm Account Deletion</title>
	</head>
	<body style="font-family: sans-serif; background: #f4f4f4; padding: 20px;">
		<div style="max-width: 600px; margin: auto; background: white; padding: 20px; border-radius: 8px;">
			<h2 style="color: #333;">Confirm Account Deletion</h2>
			<p>This is a request to delete your account on <strong>%s</strong>.</p>
			<p>If you made this request, please confirm by clicking the button below:</p>
			<p style="text-align: center; margin: 24px 0;">
				<a href="%s" style="background: #e11d48; color: white; padding: 12px 24px; text-decoration: none; border-radius: 4px;">
					Delete Account
				</a>
			</p>
			<p>If the button doesn't work, paste this link in your browser:</p>
			<p><a href="%s">%s</a></p>
			<p>If you did not request account deletion, please contact our support team at 
			<a href="mailto:%s">%s</a>.</p>
			<hr style="margin: 24px 0; border: none; border-top: 1px solid #ddd;">
			<p style="font-size: 12px; color: #666;">This message was sent automatically. Please do not reply directly to this email.</p>
		</div>
	</body>
	</html>
	`, sysName, verifyLink, verifyLink, verifyLink, support, support)
}
