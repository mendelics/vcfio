package vcfio

import (
	"fmt"
	"log"
	"os"
)

func CheckVcfSamples(vcfFiles []string) ([]string, error) {
	// Each sample gets their own variant file, their own coverage+minivars, and their own gene-variant map
	var samples []string

	missingVcfs := checkIfVcfsExist(vcfFiles)
	if len(missingVcfs) != 0 {
		return nil, fmt.Errorf("missing vcfs: %v", missingVcfs)
	}

	for i, vcf := range vcfFiles {
		// Read VCF file into query stream
		vcfReader, err := ReadVcf(vcf)
		if err != nil {
			return nil, fmt.Errorf("error reading vcf: %v", missingVcfs)
		}

		// Check if samples match in vcfs
		if i != 0 {
			checkSampleMatch(samples, vcfReader.Header.SampleNames, vcf)
		}
		samples = vcfReader.Header.SampleNames
	}

	return samples, nil
}

func checkIfVcfsExist(vcfFiles []string) []string {
	missingVcfs := make([]string, 0)
	for _, vcf := range vcfFiles {
		if !fileExists(vcf) {
			missingVcfs = append(missingVcfs, vcf)
		}
	}
	return missingVcfs
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func checkSampleMatch(sampleNames, samplesFromFirstVcf []string, vcf string) {
	if len(sampleNames) != len(samplesFromFirstVcf) {
		log.Fatalf("Number of samples in first vcf %d don't match %d in vcf %s\n", len(samplesFromFirstVcf), len(sampleNames), vcf)
	}
	for j := range sampleNames {
		if sampleNames[j] != samplesFromFirstVcf[j] {
			log.Fatalf("Sample %s in first vcf doesn't match sample %s in vcf %s\n", samplesFromFirstVcf[j], sampleNames[j], vcf)
		}
	}
}
