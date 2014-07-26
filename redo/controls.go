// 7 july 2014

package ui

// Control represents a control.
// All Controls have event handlers that take a single argument (the Doer active during the event) and return nothing.
type Control interface {
	controlParent			// platform-specific
	// TODO enable/disable (public)
	// TODO show/hide (public)
	containerShow()		// for Windows, where all controls need ot belong to an overlapped window, not to a container control; these respect programmer settings
	containerHide()
	controlSizing
}

// Button is a clickable button that performs some task.
type Button interface {
	Control

	// OnClicked sets the event handler for when the Button is clicked.
	OnClicked(func())

	// Text and SetText get and set the Button's label text.
	Text() string
	SetText(text string)
}

// NewButton creates a new Button with the given label text.
func NewButton(text string) Button {
	return newButton(text)
}

// Checkbox is a clickable box that indicates some Boolean value.
type Checkbox interface {
	Control

	// OnClicked sets the event handler for when the Checkbox is clicked (to change its toggle state).
	// TODO change to OnToggled() or OnChecked()?
	OnClicked(func())

	// Text and SetText are Requests that get and set the Checkbox's label text.
	Text() string
	SetText(text string)

	// Checked and SetChecked get and set the Checkbox's check state.
	Checked() bool
	SetChecked(checked bool)
}

// NewCheckbox creates a new Checkbox with the given label text.
// The Checkbox will be initially unchecked.
func NewCheckbox(text string) Checkbox {
	return newCheckbox(text)
}

// TextField is a Control in which the user can enter a single line of text.
type TextField interface {
	Control

	// TODO figure out what events are appropriate

	// Text and SetText are Requests that get and set the TextField's text.
	Text() string
	SetText(text string)
}

// NewTextField creates a new TextField.
func NewTextField() TextField {
	return newTextField()
}

// NewPasswordField creates a new TextField for entering passwords; that is, it hides the text being entered.
func NewPasswordField() TextField {
	return newPasswordField()
}
