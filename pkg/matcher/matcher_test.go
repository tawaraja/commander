package matcher

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_NewMatcher_Text(t *testing.T) {
    got := NewMatcher(Text)
    assert.IsType(t, TextMatcher{}, got)
}

func Test_NewMatcher_Contains(t *testing.T) {
    got := NewMatcher(Contains)
    assert.IsType(t, ContainsMatcher{}, got)
}

func Test_NewMatcher_Equal(t *testing.T) {
    got := NewMatcher(Equal)
    assert.IsType(t, EqualMatcher{}, got)
}

func TestTextMatcher_Validate(t *testing.T) {
    m := TextMatcher{}
    got := m.Match("test", "test")
    assert.True(t, got.Success)
}

func TestTextMatcher_ValidateFails(t *testing.T) {
    m := TextMatcher{}
    got := m.Match("test", "unequal")
    assert.False(t, got.Success)
    assert.Contains(t, got.Diff, "+unequal")
    assert.Contains(t, got.Diff, "-test")
}

func TestEqualMatcher_Validate(t *testing.T) {
    m := EqualMatcher{}
    got := m.Match(1, 1)
    assert.True(t, got.Success)
}

func TestEqualMatcher_ValidateFails(t *testing.T) {
    m := EqualMatcher{}
    got := m.Match(1, 0)
    assert.False(t, got.Success)
    assert.Contains(t, got.Diff, "+0")
    assert.Contains(t, got.Diff, "-1")
}
