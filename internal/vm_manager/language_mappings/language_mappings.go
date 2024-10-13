package language_mappings

import "fmt"

// LanguageKey represents the available programming languages
type LanguageKey string

const (
	Bash5_0_0             LanguageKey = "bash5.0.0"
	C_Clang7_0_1          LanguageKey = "c_clang7.0.1"
	C_Clang9_0_1          LanguageKey = "c_clang9.0.1"
	C_Clang10_0_1         LanguageKey = "c_clang10.0.1"
	C_GCC7_4_0            LanguageKey = "c_gcc7.4.0"
	C_GCC8_3_0            LanguageKey = "c_gcc8.3.0"
	C_GCC9_2_0            LanguageKey = "c_gcc9.2.0"
	CSharp_Mono6_6_0_161  LanguageKey = "csharp_mono6.6.0.161"
	CSharp_Mono6_10_0_104 LanguageKey = "csharp_mono6.10.0.104"
	CSharp_DotNet3_1_202  LanguageKey = "csharp_dotnet3.1.202"
	CSharp_DotNet3_1_302  LanguageKey = "csharp_dotnet3.1.302"
	CPP_Clang7_0_1        LanguageKey = "cpp_clang7.0.1"
	CPP_Clang9_0_1        LanguageKey = "cpp_clang9.0.1"
	CPP_Clang10_0_1       LanguageKey = "cpp_clang10.0.1"
	CPP_GCC7_4_0          LanguageKey = "cpp_gcc7.4.0"
	CPP_GCC8_3_0          LanguageKey = "cpp_gcc8.3.0"
	CPP_GCC9_2_0          LanguageKey = "cpp_gcc9.2.0"
	Clojure1_10_1         LanguageKey = "clojure1.10.1"
	CommonLisp2_0_0       LanguageKey = "commonlisp2.0.0"
	D2_089_1              LanguageKey = "d2.089.1"
	Elixir1_9_4           LanguageKey = "elixir1.9.4"
	Erlang22_2            LanguageKey = "erlang22.2"
	Executable            LanguageKey = "executable"
	FSharp_DotNet3_1_202  LanguageKey = "fsharp_dotnet3.1.202"
	FSharp_DotNet3_1_302  LanguageKey = "fsharp_dotnet3.1.302"
	Fortran9_2_0          LanguageKey = "fortran9.2.0"
	Go1_13_5              LanguageKey = "go1.13.5"
	Groovy3_0_3           LanguageKey = "groovy3.0.3"
	Haskell8_8_1          LanguageKey = "haskell8.8.1"
	Java13_0_1            LanguageKey = "java13.0.1"
	Java14_0_1            LanguageKey = "java14.0.1"
	JavaScript12_14_0     LanguageKey = "javascript12.14.0"
	Kotlin1_3_70          LanguageKey = "kotlin1.3.70"
	Lua5_3_5              LanguageKey = "lua5.3.5"
	NimStable             LanguageKey = "nimstable"
	OCaml4_09_0           LanguageKey = "ocaml4.09.0"
	Octave5_1_0           LanguageKey = "octave5.1.0"
	Perl5_28_1            LanguageKey = "perl5.28.1"
	PHP7_4_1              LanguageKey = "php7.4.1"
	PlainText             LanguageKey = "plaintext"
	Prolog1_4_5           LanguageKey = "prolog1.4.5"
	Python2_7_17          LanguageKey = "python2.7.17"
	Python3_8_1           LanguageKey = "python3.8.1"
	R4_0_0                LanguageKey = "r4.0.0"
	Ruby2_7_0             LanguageKey = "ruby2.7.0"
	Rust1_40_0            LanguageKey = "rust1.40.0"
	Scala2_13_2           LanguageKey = "scala2.13.2"
	Swift5_2_3            LanguageKey = "swift5.2.3"
	TypeScript3_7_4       LanguageKey = "typescript3.7.4"
	VBNet0_0_0_5943       LanguageKey = "vbnet0.0.0.5943"
)

func ValidateLanguageKey(key string) (LanguageKey, error) {
	validKey := LanguageKey(key)

	// Check if the key exists in the LanguageToDockerImage map
	if _, ok := LanguageToDockerImage[validKey]; !ok {
		return "", fmt.Errorf("invalid language key: %s", key)
	}

	return validKey, nil
}

var LanguageDetail = map[LanguageKey]string{
	Bash5_0_0:             "Bash (5.0.0)",
	C_Clang7_0_1:          "C (Clang 7.0.1)",
	C_Clang9_0_1:          "C (Clang 9.0.1)",
	C_Clang10_0_1:         "C (Clang 10.0.1)",
	C_GCC7_4_0:            "C (GCC 7.4.0)",
	C_GCC8_3_0:            "C (GCC 8.3.0)",
	C_GCC9_2_0:            "C (GCC 9.2.0)",
	CSharp_Mono6_6_0_161:  "C# (Mono 6.6.0.161)",
	CSharp_Mono6_10_0_104: "C# (Mono 6.10.0.104)",
	CSharp_DotNet3_1_202:  "C# (.NET Core SDK 3.1.202)",
	CSharp_DotNet3_1_302:  "C# (.NET Core SDK 3.1.302)",
	CPP_Clang7_0_1:        "C++ (Clang 7.0.1)",
	CPP_Clang9_0_1:        "C++ (Clang 9.0.1)",
	CPP_Clang10_0_1:       "C++ (Clang 10.0.1)",
	CPP_GCC7_4_0:          "C++ (GCC 7.4.0)",
	CPP_GCC8_3_0:          "C++ (GCC 8.3.0)",
	CPP_GCC9_2_0:          "C++ (GCC 9.2.0)",
	Clojure1_10_1:         "Clojure (1.10.1)",
	CommonLisp2_0_0:       "Common Lisp (SBCL 2.0.0)",
	D2_089_1:              "D (DMD 2.089.1)",
	Elixir1_9_4:           "Elixir (1.9.4)",
	Erlang22_2:            "Erlang (OTP 22.2)",
	Executable:            "Executable",
	FSharp_DotNet3_1_202:  "F# (.NET Core SDK 3.1.202)",
	FSharp_DotNet3_1_302:  "F# (.NET Core SDK 3.1.302)",
	Fortran9_2_0:          "Fortran (GFortran 9.2.0)",
	Go1_13_5:              "Go (1.13.5)",
	Groovy3_0_3:           "Groovy (3.0.3)",
	Haskell8_8_1:          "Haskell (GHC 8.8.1)",
	Java13_0_1:            "Java (OpenJDK 13.0.1)",
	Java14_0_1:            "Java (OpenJDK 14.0.1)",
	JavaScript12_14_0:     "JavaScript (Node.js 12.14.0)",
	Kotlin1_3_70:          "Kotlin (1.3.70)",
	Lua5_3_5:              "Lua (5.3.5)",
	NimStable:             "Nim (stable)",
	OCaml4_09_0:           "OCaml (4.09.0)",
	Octave5_1_0:           "Octave (5.1.0)",
	Perl5_28_1:            "Perl (5.28.1)",
	PHP7_4_1:              "PHP (7.4.1)",
	PlainText:             "Plain Text",
	Prolog1_4_5:           "Prolog (GNU Prolog 1.4.5)",
	Python2_7_17:          "Python (2.7.17)",
	Python3_8_1:           "Python (3.8.1)",
	R4_0_0:                "R (4.0.0)",
	Ruby2_7_0:             "Ruby (2.7.0)",
	Rust1_40_0:            "Rust (1.40.0)",
	Scala2_13_2:           "Scala (2.13.2)",
	Swift5_2_3:            "Swift (5.2.3)",
	TypeScript3_7_4:       "TypeScript (3.7.4)",
	VBNet0_0_0_5943:       "Visual Basic.Net (vbnc 0.0.0.5943)",
}

var LanguageToDockerImage = map[LanguageKey]string{
	Bash5_0_0:             "docker.io/library/bash:5.0",
	C_Clang7_0_1:          "docker.io/library/buildpack-deps:buster",
	C_Clang9_0_1:          "docker.io/library/buildpack-deps:buster",
	C_Clang10_0_1:         "docker.io/library/buildpack-deps:buster",
	C_GCC7_4_0:            "docker.io/library/gcc:7.4",
	C_GCC8_3_0:            "docker.io/library/gcc:8.3",
	C_GCC9_2_0:            "docker.io/library/gcc:9.2",
	CSharp_Mono6_6_0_161:  "docker.io/library/mono:6.6.0.161",
	CSharp_Mono6_10_0_104: "docker.io/library/mono:6.10.0.104",
	CSharp_DotNet3_1_202:  "mcr.microsoft.com/dotnet/core/sdk:3.1.202",
	CSharp_DotNet3_1_302:  "mcr.microsoft.com/dotnet/core/sdk:3.1.302",
	CPP_Clang7_0_1:        "docker.io/library/buildpack-deps:buster",
	CPP_Clang9_0_1:        "docker.io/library/buildpack-deps:buster",
	CPP_Clang10_0_1:       "docker.io/library/buildpack-deps:buster",
	CPP_GCC7_4_0:          "docker.io/library/gcc:7.4",
	CPP_GCC8_3_0:          "docker.io/library/gcc:8.3",
	CPP_GCC9_2_0:          "docker.io/library/gcc:9.2",
	Clojure1_10_1:         "docker.io/library/clojure:openjdk-8-lein-2.9.1",
	CommonLisp2_0_0:       "docker.io/library/sbcl:2.0.0",
	D2_089_1:              "docker.io/library/dlang2/dmd-ubuntu:2.089.1",
	Elixir1_9_4:           "docker.io/library/elixir:1.9.4",
	Erlang22_2:            "docker.io/library/erlang:22.2",
	Executable:            "docker.io/library/ubuntu:20.04",
	FSharp_DotNet3_1_202:  "mcr.microsoft.com/dotnet/core/sdk:3.1.202",
	FSharp_DotNet3_1_302:  "mcr.microsoft.com/dotnet/core/sdk:3.1.302",
	Fortran9_2_0:          "docker.io/library/gcc:9.2",
	Go1_13_5:              "docker.io/library/golang:1.13.5",
	Groovy3_0_3:           "docker.io/library/groovy:3.0.3-jdk8",
	Haskell8_8_1:          "docker.io/library/haskell:8.8.1",
	Java13_0_1:            "docker.io/library/openjdk:13.0.1",
	Java14_0_1:            "docker.io/library/openjdk:14.0.1",
	JavaScript12_14_0:     "docker.io/library/node:12.14.0",
	Kotlin1_3_70:          "docker.io/library/openjdk:8-jdk",
	Lua5_3_5:              "docker.io/library/lua:5.3",
	NimStable:             "docker.io/library/nimlang/nim:latest",
	OCaml4_09_0:           "docker.io/library/ocaml:4.09",
	Octave5_1_0:           "docker.io/library/gnuoctave:5.1.0",
	Perl5_28_1:            "docker.io/library/perl:5.28.1",
	PHP7_4_1:              "docker.io/library/php:7.4.1",
	PlainText:             "docker.io/library/ubuntu:20.04",
	Prolog1_4_5:           "docker.io/library/swipl:latest",
	Python2_7_17:          "docker.io/library/python:2.7.17",
	Python3_8_1:           "docker.io/library/python:3.8.1",
	R4_0_0:                "docker.io/library/r-base:4.0.0",
	Ruby2_7_0:             "docker.io/library/ruby:2.7.0",
	Rust1_40_0:            "docker.io/library/rust:1.40.0",
	Scala2_13_2:           "docker.io/library/scala:2.13.2",
	Swift5_2_3:            "docker.io/library/swift:5.2.3",
	TypeScript3_7_4:       "docker.io/library/node:12.14.0",
	VBNet0_0_0_5943:       "mcr.microsoft.com/dotnet/core/sdk:3.1",
}

var PendingBuildMap = map[string]string{
	"Assembly (NASM 2.14.02)":                       "docker.io/library/ubuntu:20.04",
	"Basic (FBC 1.07.1)":                            "docker.io/library/ubuntu:20.04",
	"Bosque (latest)":                               "docker.io/library/ubuntu:20.04",
	"C# Test (.NET Core SDK 3.1.302, NUnit 3.12.0)": "mcr.microsoft.com/dotnet/core/sdk:3.1.302",
	"C++ Test (Clang 10.0.1, Google Test 1.8.1)":    "docker.io/library/buildpack-deps:buster",
	"C++ Test (GCC 8.4.0, Google Test 1.8.1)":       "docker.io/library/gcc:8.4",
	"C3 (latest)":                                   "docker.io/library/ubuntu:20.04",
	"COBOL (GnuCOBOL 2.2)":                          "docker.io/library/ubuntu:20.04",
	"Java Test (OpenJDK 14.0.1, JUnit Platform Console Standalone 1.6.2)": "docker.io/library/openjdk:14.0.1",
	"MPI (OpenRTE 3.1.3) with C (GCC 8.4.0)":                              "docker.io/library/ubuntu:20.04",
	"MPI (OpenRTE 3.1.3) with C++ (GCC 8.4.0)":                            "docker.io/library/ubuntu:20.04",
	"MPI (OpenRTE 3.1.3) with Python (3.7.7)":                             "docker.io/library/python:3.7.7",
	"Objective-C (Clang 7.0.1)":                                           "docker.io/library/ubuntu:20.04",
	"Pascal (FPC 3.0.4)":                                                  "docker.io/library/ubuntu:20.04",
	"Python for ML (3.7.7)":                                               "docker.io/library/python:3.7.7",
	"SQL (SQLite 3.27.2)":                                                 "docker.io/library/ubuntu:20.04",
}
