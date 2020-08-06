// Code generated by https://github.com/gagliardetto/codebox. DO NOT EDIT.

package main

import (
	"archive/zip"
	"io"
	"os"
)

func TaintStepTest_ArchiveZipFileInfoHeader_B0I0O0(sourceCQL interface{}) interface{} {
	fromFileInfo656 := sourceCQL.(os.FileInfo)
	intoFileHeader414, _ := zip.FileInfoHeader(fromFileInfo656)
	return intoFileHeader414
}

func TaintStepTest_ArchiveZipNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	fromReaderAt518 := sourceCQL.(io.ReaderAt)
	intoReader650, _ := zip.NewReader(fromReaderAt518, 0)
	return intoReader650
}

func TaintStepTest_ArchiveZipNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	fromWriter784 := sourceCQL.(*zip.Writer)
	var intoWriter957 io.Writer
	intermediateCQL := zip.NewWriter(intoWriter957)
	link(fromWriter784, intermediateCQL)
	return intoWriter957
}

func TaintStepTest_ArchiveZipOpenReader_B0I0O0(sourceCQL interface{}) interface{} {
	fromString520 := sourceCQL.(string)
	intoReadCloser443, _ := zip.OpenReader(fromString520)
	return intoReadCloser443
}

func TaintStepTest_ArchiveZipFileOpen_B0I0O0(sourceCQL interface{}) interface{} {
	fromFile127 := sourceCQL.(zip.File)
	intoReadCloser483, _ := fromFile127.Open()
	return intoReadCloser483
}

func TaintStepTest_ArchiveZipWriterCreate_B0I0O0(sourceCQL interface{}) interface{} {
	fromWriter989 := sourceCQL.(io.Writer)
	var intoWriter982 zip.Writer
	intermediateCQL, _ := intoWriter982.Create("")
	link(fromWriter989, intermediateCQL)
	return intoWriter982
}

func TaintStepTest_ArchiveZipWriterCreateHeader_B0I0O0(sourceCQL interface{}) interface{} {
	fromWriter417 := sourceCQL.(io.Writer)
	var intoWriter584 zip.Writer
	intermediateCQL, _ := intoWriter584.CreateHeader(nil)
	link(fromWriter417, intermediateCQL)
	return intoWriter584
}

func RunAllTaints_ArchiveZip() {
	{
		source := newSource(0)
		out := TaintStepTest_ArchiveZipFileInfoHeader_B0I0O0(source)
		sink(0, out)
	}
	{
		source := newSource(1)
		out := TaintStepTest_ArchiveZipNewReader_B0I0O0(source)
		sink(1, out)
	}
	{
		source := newSource(2)
		out := TaintStepTest_ArchiveZipNewWriter_B0I0O0(source)
		sink(2, out)
	}
	{
		source := newSource(3)
		out := TaintStepTest_ArchiveZipOpenReader_B0I0O0(source)
		sink(3, out)
	}
	{
		source := newSource(4)
		out := TaintStepTest_ArchiveZipFileOpen_B0I0O0(source)
		sink(4, out)
	}
	{
		source := newSource(5)
		out := TaintStepTest_ArchiveZipWriterCreate_B0I0O0(source)
		sink(5, out)
	}
	{
		source := newSource(6)
		out := TaintStepTest_ArchiveZipWriterCreateHeader_B0I0O0(source)
		sink(6, out)
	}
}