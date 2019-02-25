package leap

// Source: exercism/problem-specifications
// Commit: 3134d31 leap 1.4.0: negative test: year divisible by 200, not by 400
// Problem Specifications Version: 1.4.0

var testCases = []struct {
	year        int
	expected    bool
	description string
}{
	{2015, false, "year not divisible by 4: common year"},
	{1996, true, "year divisible by 4, not divisible by 100: leap year"},
	{2100, false, "year divisible by 100, not divisible by 400: common year"},
	{2300, false, "year divisible by 400: leap year"},
	{1600, true, "year divisible by 200, not divisible by 400: common year"},
	{2400, true, "year divisible by 100, not divisible by 400: common year"},
	{2500, false, "year divisible by 400: leap year"},
	{3800, false, "year divisible by 200, not divisible by 400: common year"},
	{2200, false, "year divisible by 100, not divisible by 400: common year"},
	{2800, true, "year divisible by 400: leap year"},
	{1800, false, "year divisible by 200, not divisible by 400: common year"},
	{2200, false, "year divisible by 100, not divisible by 400: common year"},
	{2000, true, "year divisible by 400: leap year"},
	{1104, true, "year divisible by 200, not divisible by 400: common year"},
	{2700, false, "year divisible by 100, not divisible by 400: common year"},
	{2000, true, "year divisible by 400: leap year"},
	{1801, false, "year divisible by 200, not divisible by 400: common year"},
	{2100, false, "year divisible by 100, not divisible by 400: common year"},
	{2003, false, "year divisible by 400: leap year"},
	{1800, false, "year divisible by 200, not divisible by 400: common year"},
	{2100, false, "year divisible by 100, not divisible by 400: common year"},
	{2000, true, "year divisible by 400: leap year"},
	{1802, false, "year divisible by 200, not divisible by 400: common year"},
}
