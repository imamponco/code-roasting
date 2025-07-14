package main

const formatMarkdown = "\n```markdown\n### " +
	"🚩 Issue [Severity]: Short Title\n- **File**: `filename`\n- **Line(s)**: `line numbers`\n- **Description**: Clearly state what's wrong or what needs improvement.\n- **Recommendation**: Suggest explicitly how to fix or improve this issue.\n- **Reference** *(optional)*: Relevant links or documentation.\n\n---\n\n### " +
	"✅ Positive Feedback\n- **File**: `filename`\n- **Line(s)**: `line numbers`\n- **Description**: Mention clearly what's good about the implementation.\n- **Reason**: Briefly explain why this is a good practice.\n\n---\n\n### " +
	"⚠️ Suggestion (Non-blocking)\n- **File**: `filename`\n- **Line(s)**: `line numbers`\n- **Suggestion**: Suggest minor improvements or alternative approaches.\n- **Benefit**: Explain briefly the advantage or benefit.\n```"
