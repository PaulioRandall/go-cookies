package cookies

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripSpace(t *testing.T) {
	require.Equal(t, "Rincewind", StripSpace("Rince \n\t\f\r wind"))
	require.Equal(t, "Rincewind", StripSpace("\t \n\t \r\n\n\fRincewind"))
	require.Equal(t, "Rincewind", StripSpace("Rincewind\r\n \t\t\f \r  \v\v"))
	require.Equal(t, "Rincewind", StripSpace("\r\nRi \tn\tc\t\t ew\f \r in\vd\v"))
	require.Equal(t, "Rincewind", StripSpace("Rincewind"))
	require.Equal(t, "", StripSpace(""))
	require.Equal(t, "", StripSpace("\r\n \t\t \t\t \f \r  \v\v  "))
}

func TestIndentLines(t *testing.T) {
	require.Equal(t,
		"\t\t\n\t\tMoonglow\n\t\tMoonglow\n\t\t",
		IndentLines(2, "\t", "\nMoonglow\nMoonglow\n"))
	require.Equal(t, "Moonglow", IndentLines(1, "", "Moonglow"))
	require.Equal(t, "Moonglow", IndentLines(0, "\t", "Moonglow"))
	require.Equal(t, "\t", IndentLines(1, "\t", ""))
	require.Panics(t, func() {
		IndentLines(-5, "\t", "Moonglow")
	})
}
