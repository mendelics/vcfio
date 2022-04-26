package vcfio

import (
	"fmt"
	"log"
)

func Example_vcfReader() {
	vcfScanner, header, err := ReadNewVcf("samples/cnv.vcf.gz")
	if err != nil {
		log.Fatalf("Error reading vcf, %v\n", err)
	}

	for vcfScanner.Scan() {
		line := vcfScanner.Text()
		variantInfo, quality, genotypes, mafs := ParseVariant(line, header)
		fmt.Println(variantInfo.VariantKey, quality.QualScore, genotypes, mafs.Gnomad)
	}
	// Output:
	//2-165959843-166169412-DEL 23 [{TEST-001 1 false  0 0 false false false false}] {0 0 0 0 0 0 0 0 0 }
	//7-117469210-117679939-DEL 23 [{TEST-001 2 false  0 0 false false false false}] {0 0 0 0 0 0 0 0 0 }
	//17-15229777-15265357-DUP 9.97 [{TEST-001 1 false  0 0 false false false false}] {0 0 0 0 0 0 0 0 0 }
	//X-154021013-154101781-DEL 11.7 [{TEST-001 1 false  0 0 false false false false}] {0 0 0 0 0 0 0 0 0 }
}
