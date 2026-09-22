package main

// Question represents a single quiz question.
type Question struct {
	ID           int
	Category     string
	Difficulty   string // Easy, Medium, Hard
	Text         string
	Options      [4]string
	CorrectIndex int // 0-3
	Explanation  string
}

// QuestionBank holds all available questions.
var QuestionBank = []Question{
	{
		ID: 1, Category: "Phishing", Difficulty: "Easy",
		Text: "You receive an email claiming to be from your bank, urgently asking you to click a link to 'verify your account' or it will be suspended. What should you do?",
		Options: [4]string{
			"Click the link immediately to avoid losing access",
			"Reply with your account details to be safe",
			"Go directly to the bank's official website or app instead of clicking the link",
			"Forward it to friends to warn them",
		},
		CorrectIndex: 2,
		Explanation:  "Urgency and threats are classic phishing tactics. Always navigate to the official site directly rather than clicking embedded links.",
	},
	{
		ID: 2, Category: "Passwords", Difficulty: "Easy",
		Text: "Which of the following is the strongest password?",
		Options: [4]string{
			"password123",
			"Tr0ub4dor&3",
			"correct-horse-battery-staple-92!",
			"qwerty",
		},
		CorrectIndex: 2,
		Explanation:  "Long passphrases with random words, numbers, and symbols are harder to crack than short complex-looking ones.",
	},
	{
		ID: 3, Category: "Passwords", Difficulty: "Medium",
		Text: "What is the main benefit of using a password manager?",
		Options: [4]string{
			"It makes your passwords visible to websites for convenience",
			"It lets you use one weak password everywhere",
			"It generates and stores unique, strong passwords for every account",
			"It removes the need for two-factor authentication",
		},
		CorrectIndex: 2,
		Explanation:  "Password managers let you use a unique, strong password per site without needing to memorize them all.",
	},
	{
		ID: 4, Category: "Social Engineering", Difficulty: "Medium",
		Text: "A caller claims to be from IT support and asks for your password to 'fix an issue'. What should you do?",
		Options: [4]string{
			"Give them your password since they say they're from IT",
			"Refuse and verify their identity through official channels first",
			"Give them a slightly different password",
			"Ask a coworker to give their password instead",
		},
		CorrectIndex: 1,
		Explanation:  "Legitimate IT staff never need your actual password. This is a classic pretexting/social engineering attack.",
	},
	{
		ID: 5, Category: "Malware", Difficulty: "Easy",
		Text: "What is 'ransomware'?",
		Options: [4]string{
			"Software that shows you ads",
			"Malware that encrypts your files and demands payment to unlock them",
			"A tool to recover deleted files",
			"A type of antivirus software",
		},
		CorrectIndex: 1,
		Explanation:  "Ransomware encrypts victims' data and demands a ransom, usually in cryptocurrency, for the decryption key.",
	},
	{
		ID: 6, Category: "Safe Browsing", Difficulty: "Easy",
		Text: "You want to check if a website is using an encrypted connection. What should you look for?",
		Options: [4]string{
			"A padlock icon and 'https://' in the address bar",
			"A colorful logo on the homepage",
			"The word 'secure' somewhere on the page",
			"A pop-up saying the site is safe",
		},
		CorrectIndex: 0,
		Explanation:  "HTTPS with a padlock icon indicates an encrypted connection, though it doesn't guarantee the site itself is trustworthy.",
	},
	{
		ID: 7, Category: "Data Privacy", Difficulty: "Medium",
		Text: "Before installing a new mobile app, what is a good security practice?",
		Options: [4]string{
			"Accept all permissions without reading them to save time",
			"Review the permissions it requests and whether they make sense for its function",
			"Only check the star rating",
			"Install it and worry about permissions later",
		},
		CorrectIndex: 1,
		Explanation:  "Apps should only request permissions relevant to their function. A flashlight app asking for contacts access is a red flag.",
	},
	{
		ID: 8, Category: "Wi-Fi Security", Difficulty: "Medium",
		Text: "You're at a coffee shop and need to check your bank account online. What's the safest approach?",
		Options: [4]string{
			"Use the shop's free public Wi-Fi directly",
			"Use your phone's mobile data or a trusted VPN instead of public Wi-Fi",
			"Ask the barista for the Wi-Fi password and proceed normally",
			"Use public Wi-Fi but browse in an incognito window",
		},
		CorrectIndex: 1,
		Explanation:  "Public Wi-Fi can be intercepted or spoofed. Mobile data or a reputable VPN is much safer for sensitive activity.",
	},
	{
		ID: 9, Category: "Phishing", Difficulty: "Medium",
		Text: "Which of these is a common red flag in a phishing email?",
		Options: [4]string{
			"Correct spelling and grammar throughout",
			"An email from a colleague you email daily",
			"Mismatched or suspicious sender addresses and urgent, threatening language",
			"An email with no links or attachments",
		},
		CorrectIndex: 2,
		Explanation:  "Phishing emails often use spoofed or slightly-off sender addresses combined with urgency to pressure quick action.",
	},
	{
		ID: 10, Category: "Authentication", Difficulty: "Easy",
		Text: "What does Two-Factor Authentication (2FA) add to your login process?",
		Options: [4]string{
			"A second password that's the same as the first",
			"An additional verification step, like a code from your phone, beyond just a password",
			"Nothing, it's just a formality",
			"A requirement to change your password every login",
		},
		CorrectIndex: 1,
		Explanation:  "2FA requires something you know (password) plus something you have (phone/token) or are (biometric), greatly reducing account takeover risk.",
	},
}
