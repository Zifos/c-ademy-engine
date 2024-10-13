package commands

import (
	"c-ademy/internal/vm_manager/language_mappings"
	"fmt"
	"path/filepath"
	"strings"
)

// GenerateDockerCommand creates a Docker CMD array for the given language and file
func GenerateDockerCommand(lang language_mappings.LanguageKey, filePath string, extraArgs ...string) ([]string, error) {
	_, ok := language_mappings.LanguageToDockerImage[lang]
	if !ok {
		return nil, fmt.Errorf("unsupported language: %s", lang)
	}

	var cmd []string
	fileName := filepath.Base(filePath)
	fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	switch lang {
	case language_mappings.Bash5_0_0:
		cmd = []string{"bash", filePath}
	case language_mappings.C_Clang7_0_1, language_mappings.C_Clang9_0_1, language_mappings.C_Clang10_0_1,
		language_mappings.C_GCC7_4_0, language_mappings.C_GCC8_3_0, language_mappings.C_GCC9_2_0:
		cmd = []string{"sh", "-c", fmt.Sprintf("gcc %s -o /tmp/out && /tmp/out", filePath)}
	case language_mappings.CSharp_Mono6_6_0_161, language_mappings.CSharp_Mono6_10_0_104:
		cmd = []string{"mcs", "-out:/tmp/out.exe", filePath, "&&", "mono", "/tmp/out.exe"}
	case language_mappings.CSharp_DotNet3_1_202, language_mappings.CSharp_DotNet3_1_302:
		cmd = []string{"dotnet", "run", "--project", filePath}
	case language_mappings.CPP_Clang7_0_1, language_mappings.CPP_Clang9_0_1, language_mappings.CPP_Clang10_0_1,
		language_mappings.CPP_GCC7_4_0, language_mappings.CPP_GCC8_3_0, language_mappings.CPP_GCC9_2_0:
		cmd = []string{"sh", "-c", fmt.Sprintf("g++ %s -o /tmp/out && /tmp/out", filePath)}
	case language_mappings.Clojure1_10_1:
		cmd = []string{"clojure", filePath}
	case language_mappings.CommonLisp2_0_0:
		cmd = []string{"sbcl", "--script", filePath}
	case language_mappings.D2_089_1:
		cmd = []string{"sh", "-c", fmt.Sprintf("dmd %s -of/tmp/out && /tmp/out", filePath)}
	case language_mappings.Elixir1_9_4:
		cmd = []string{"elixir", filePath}
	case language_mappings.Erlang22_2:
		cmd = []string{"escript", filePath}
	case language_mappings.Executable:
		cmd = []string{filePath}
	case language_mappings.FSharp_DotNet3_1_202, language_mappings.FSharp_DotNet3_1_302:
		cmd = []string{"dotnet", "fsi", filePath}
	case language_mappings.Fortran9_2_0:
		cmd = []string{"sh", "-c", fmt.Sprintf("gfortran %s -o /tmp/out && /tmp/out", filePath)}
	case language_mappings.Go1_13_5:
		cmd = []string{"go", "run", filePath}
	case language_mappings.Groovy3_0_3:
		cmd = []string{"groovy", filePath}
	case language_mappings.Haskell8_8_1:
		cmd = []string{"runhaskell", filePath}
	case language_mappings.Java13_0_1, language_mappings.Java14_0_1:
		cmd = []string{"sh", "-c", fmt.Sprintf("javac %s && java %s", fileName, fileNameWithoutExt)}
	case language_mappings.JavaScript12_14_0:
		cmd = []string{"node", filePath}
	case language_mappings.Kotlin1_3_70:
		cmd = []string{"sh", "-c", fmt.Sprintf("kotlinc %s -include-runtime -d /tmp/out.jar && java -jar /tmp/out.jar", filePath)}
	case language_mappings.Lua5_3_5:
		cmd = []string{"lua", filePath}
	case language_mappings.NimStable:
		cmd = []string{"nim", "c", "-r", filePath}
	case language_mappings.OCaml4_09_0:
		cmd = []string{"ocaml", filePath}
	case language_mappings.Octave5_1_0:
		cmd = []string{"octave", filePath}
	case language_mappings.Perl5_28_1:
		cmd = []string{"perl", filePath}
	case language_mappings.PHP7_4_1:
		cmd = []string{"php", filePath}
	case language_mappings.PlainText:
		cmd = []string{"cat", filePath}
	case language_mappings.Prolog1_4_5:
		cmd = []string{"swipl", "-q", "-f", filePath, "-t", "halt"}
	case language_mappings.Python2_7_17:
		cmd = []string{"python2", filePath}
	case language_mappings.Python3_8_1:
		cmd = []string{"python3", filePath}
	case language_mappings.R4_0_0:
		cmd = []string{"Rscript", filePath}
	case language_mappings.Ruby2_7_0:
		cmd = []string{"ruby", filePath}
	case language_mappings.Rust1_40_0:
		cmd = []string{"sh", "-c", fmt.Sprintf("rustc %s -o /tmp/out && /tmp/out", filePath)}
	case language_mappings.Scala2_13_2:
		cmd = []string{"scala", filePath}
	case language_mappings.Swift5_2_3:
		cmd = []string{"swift", filePath}
	case language_mappings.TypeScript3_7_4:
		cmd = []string{"sh", "-c", fmt.Sprintf("tsc %s --outFile /tmp/out.js && node /tmp/out.js", filePath)}
	case language_mappings.VBNet0_0_0_5943:
		cmd = []string{"vbnc", "-out:/tmp/out.exe", filePath, "&&", "mono", "/tmp/out.exe"}
	default:
		return nil, fmt.Errorf("command generation not implemented for language: %s", lang)
	}

	// Append extra arguments
	cmd = append(cmd, extraArgs...)

	return cmd, nil
}

// GetDockerImage returns the Docker image for the given language
func GetDockerImage(lang language_mappings.LanguageKey) (string, error) {
	dockerImage, ok := language_mappings.LanguageToDockerImage[lang]
	if !ok {
		return "", fmt.Errorf("unsupported language: %s", lang)
	}
	return dockerImage, nil
}
