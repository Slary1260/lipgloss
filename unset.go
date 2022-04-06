package lipgloss

// UnsetBold removes the bold style rule, if set.
func (s Style) UnsetBold() Style {
	s.rules.Delete(boldKey)
	return s
}

// UnsetItalic removes the italic style rule, if set.
func (s Style) UnsetItalic() Style {
	s.rules.Delete(italicKey)
	return s
}

// UnsetUnderline removes the underline style rule, if set.
func (s Style) UnsetUnderline() Style {
	s.rules.Delete(underlineKey)
	return s
}

// UnsetStrikethrough removes the strikethrough style rule, if set.
func (s Style) UnsetStrikethrough() Style {
	s.rules.Delete(strikethroughKey)
	return s
}

// UnsetReverse removes the reverse style rule, if set.
func (s Style) UnsetReverse() Style {
	s.rules.Delete(reverseKey)
	return s
}

// UnsetBlink removes the bold style rule, if set.
func (s Style) UnsetBlink() Style {
	s.rules.Delete(blinkKey)
	return s
}

// UnsetFaint removes the faint style rule, if set.
func (s Style) UnsetFaint() Style {
	s.rules.Delete(faintKey)
	return s
}

// UnsetForeground removes the foreground style rule, if set.
func (s Style) UnsetForeground() Style {
	s.rules.Delete(foregroundKey)
	return s
}

// UnsetBackground removes the background style rule, if set.
func (s Style) UnsetBackground() Style {
	s.rules.Delete(backgroundKey)
	return s
}

// UnsetWidth removes the width style rule, if set.
func (s Style) UnsetWidth() Style {
	s.rules.Delete(widthKey)
	return s
}

// UnsetHeight removes the height style rule, if set.
func (s Style) UnsetHeight() Style {
	s.rules.Delete(heightKey)
	return s
}

// UnsetAlign removes the text alignment style rule, if set.
func (s Style) UnsetAlign() Style {
	s.rules.Delete(alignKey)
	return s
}

// UnsetPadding removes all padding style rules.
func (s Style) UnsetPadding() Style {
	s.rules.Delete(paddingLeftKey)
	s.rules.Delete(paddingRightKey)
	s.rules.Delete(paddingTopKey)
	s.rules.Delete(paddingBottomKey)
	return s
}

// UnsetPaddingLeft removes the left padding style rule, if set.
func (s Style) UnsetPaddingLeft() Style {
	s.rules.Delete(paddingLeftKey)
	return s
}

// UnsetPaddingRight removes the right padding style rule, if set.
func (s Style) UnsetPaddingRight() Style {
	s.rules.Delete(paddingRightKey)
	return s
}

// UnsetPaddingTop removes the top padding style rule, if set.
func (s Style) UnsetPaddingTop() Style {
	s.rules.Delete(paddingTopKey)
	return s
}

// UnsetPaddingBottom removes the bottom style rule, if set.
func (s Style) UnsetPaddingBottom() Style {
	s.rules.Delete(paddingBottomKey)
	return s
}

// UnsetColorWhitespace removes the rule for coloring padding, if set.
func (s Style) UnsetColorWhitespace() Style {
	s.rules.Delete(colorWhitespaceKey)
	return s
}

// UnsetMargins removes all margin style rules.
func (s Style) UnsetMargins() Style {
	s.rules.Delete(marginLeftKey)
	s.rules.Delete(marginRightKey)
	s.rules.Delete(marginTopKey)
	s.rules.Delete(marginBottomKey)
	return s
}

// UnsetMarginLeft removes the left margin style rule, if set.
func (s Style) UnsetMarginLeft() Style {
	s.rules.Delete(marginLeftKey)
	return s
}

// UnsetMarginRight removes the right margin style rule, if set.
func (s Style) UnsetMarginRight() Style {
	s.rules.Delete(marginRightKey)
	return s
}

// UnsetMarginTop removes the top margin style rule, if set.
func (s Style) UnsetMarginTop() Style {
	s.rules.Delete(marginTopKey)
	return s
}

// UnsetMarginBottom removes the bottom margin style rule, if set.
func (s Style) UnsetMarginBottom() Style {
	s.rules.Delete(marginBottomKey)
	return s
}

// UnsetMarginBackground removes the margin's background color. Note that the
// margin's background color can be set from the background color of another
// style during inheritance.
func (s Style) UnsetMarginBackground() Style {
	s.rules.Delete(marginBackgroundKey)
	return s
}

// UnsetBorderStyle removes the border style rule, if set.
func (s Style) UnsetBorderStyle() Style {
	s.rules.Delete(borderStyleKey)
	return s
}

// UnsetBorderTop removes the border top style rule, if set.
func (s Style) UnsetBorderTop() Style {
	s.rules.Delete(borderTopKey)
	return s
}

// UnsetBorderRight removes the border right style rule, if set.
func (s Style) UnsetBorderRight() Style {
	s.rules.Delete(borderRightKey)
	return s
}

// UnsetBorderBottom removes the border bottom style rule, if set.
func (s Style) UnsetBorderBottom() Style {
	s.rules.Delete(borderBottomKey)
	return s
}

// UnsetBorderLeft removes the border left style rule, if set.
func (s Style) UnsetBorderLeft() Style {
	s.rules.Delete(borderLeftKey)
	return s
}

// UnsetBorderForeground removes all border foreground color styles, if set.
func (s Style) UnsetBorderForeground() Style {
	s.rules.Delete(borderTopForegroundKey)
	s.rules.Delete(borderRightForegroundKey)
	s.rules.Delete(borderBottomForegroundKey)
	s.rules.Delete(borderLeftForegroundKey)
	return s
}

// UnsetBorderTopForeground removes the top border foreground color rule,
// if set.
func (s Style) UnsetBorderTopForeground() Style {
	s.rules.Delete(borderTopForegroundKey)
	return s
}

// UnsetBorderRightForeground removes the right border foreground color rule,
// if set.
func (s Style) UnsetBorderRightForeground() Style {
	s.rules.Delete(borderRightForegroundKey)
	return s
}

// UnsetBorderBottomForeground removes the bottom border foreground color
// rule, if set.
func (s Style) UnsetBorderBottomForeground() Style {
	s.rules.Delete(borderBottomForegroundKey)
	return s
}

// UnsetBorderLeftForeground removes the left border foreground color rule,
// if set.
func (s Style) UnsetBorderLeftForeground() Style {
	s.rules.Delete(borderLeftForegroundKey)
	return s
}

// UnsetBorderBackground removes all border background color styles, if
// set.
func (s Style) UnsetBorderBackground() Style {
	s.rules.Delete(borderTopBackgroundKey)
	s.rules.Delete(borderRightBackgroundKey)
	s.rules.Delete(borderBottomBackgroundKey)
	s.rules.Delete(borderLeftBackgroundKey)
	return s
}

// UnsetBorderTopBackgroundColor removes the top border background color rule,
// if set.
func (s Style) UnsetBorderTopBackgroundColor() Style {
	s.rules.Delete(borderTopBackgroundKey)
	return s
}

// UnsetBorderRightBackground removes the right border background color
// rule, if set.
func (s Style) UnsetBorderRightBackground() Style {
	s.rules.Delete(borderRightBackgroundKey)
	return s
}

// UnsetBorderBottomBackground removes the bottom border background color
// rule, if set.
func (s Style) UnsetBorderBottomBackground() Style {
	s.rules.Delete(borderBottomBackgroundKey)
	return s
}

// UnsetBorderLeftBackground removes the left border color rule, if set.
func (s Style) UnsetBorderLeftBackground() Style {
	s.rules.Delete(borderLeftBackgroundKey)
	return s
}

// UnsetInline removes the inline style rule, if set.
func (s Style) UnsetInline() Style {
	s.rules.Delete(inlineKey)
	return s
}

// UnsetMaxWidth removes the max width style rule, if set.
func (s Style) UnsetMaxWidth() Style {
	s.rules.Delete(maxWidthKey)
	return s
}

// UnsetMaxHeight removes the max height style rule, if set.
func (s Style) UnsetMaxHeight() Style {
	s.rules.Delete(maxHeightKey)
	return s
}

// UnsetUnderlineSpaces removes the value set by UnderlineSpaces.
func (s Style) UnsetUnderlineSpaces() Style {
	s.rules.Delete(underlineSpacesKey)
	return s
}

// UnsetStrikethroughSpaces removes the value set by StrikethroughSpaces.
func (s Style) UnsetStrikethroughSpaces() Style {
	s.rules.Delete(strikethroughSpacesKey)
	return s
}

// UnsetString sets the underlying string value to the empty string.
func (s Style) UnsetString() Style {
	s.value = ""
	return s
}
