package quill

import (
	"testing"
)

func TestImageWithWidthAndStyle(t *testing.T) {
	// Test case from the user's example
	deltaJSON := `[{"insert":"ABC"},{"attributes":{"width":"348","style":"display: block; margin: 0 auto;"},"insert":{"image":"/static/uploads/questions/6041760dbe5b.jpg"}},{"insert":"\n"}]`
	
	expected := `<p>ABC<img width="348" style="display: block; margin: 0 auto;" src="/static/uploads/questions/6041760dbe5b.jpg"></p>`
	
	result, err := Render([]byte(deltaJSON))
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	resultStr := string(result)
	if resultStr != expected {
		t.Errorf("\nExpected: %s\nGot:      %s", expected, resultStr)
	}
}

func TestImageWithWidth(t *testing.T) {
	// Test with only width attribute
	deltaJSON := `[{"attributes":{"width":"500"},"insert":{"image":"/test.jpg"}},{"insert":"\n"}]`
	
	expected := `<p><img width="500" src="/test.jpg"></p>`
	
	result, err := Render([]byte(deltaJSON))
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	resultStr := string(result)
	if resultStr != expected {
		t.Errorf("\nExpected: %s\nGot:      %s", expected, resultStr)
	}
}

func TestImageWithStyle(t *testing.T) {
	// Test with only style attribute
	deltaJSON := `[{"attributes":{"style":"border: 1px solid black;"},"insert":{"image":"/test.jpg"}},{"insert":"\n"}]`
	
	expected := `<p><img style="border: 1px solid black;" src="/test.jpg"></p>`
	
	result, err := Render([]byte(deltaJSON))
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	resultStr := string(result)
	if resultStr != expected {
		t.Errorf("\nExpected: %s\nGot:      %s", expected, resultStr)
	}
}

func TestImageWithoutWidthAndStyle(t *testing.T) {
	// Test to ensure existing behavior without width and style still works
	deltaJSON := `[{"insert":{"image":"/test.jpg"}},{"insert":"\n"}]`
	
	expected := `<p><img src="/test.jpg"></p>`
	
	result, err := Render([]byte(deltaJSON))
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	resultStr := string(result)
	if resultStr != expected {
		t.Errorf("\nExpected: %s\nGot:      %s", expected, resultStr)
	}
}
