package tree_sitter_llvm_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_llvm "git+github.com/benwilliamgraham/tree-sitter-llvm.git/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_llvm.Language())
	if language == nil {
		t.Errorf("Error loading Llvm grammar")
	}
}
