package xmlsig

import "encoding/xml"

/*
Data structures to represent some of the types defined in
Schema for XML Signatures, http://www.w3.org/2000/09/xmldsig.
*/

// Signature element is the root element of an XML Signature.
type Signature struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	SignedInfo     SignedInfo
	SignatureValue string `xml:"SignatureValue"`
	KeyInfo        KeyInfo
}

// Algorithm describes the digest or signature used when digest or signature.
type Algorithm struct {
	Algorithm string `xml:",attr"`
}

// SignedInfo includes a canonicalization algorithm, a signature algorithm, and a reference.
type SignedInfo struct {
	XMLName                xml.Name  `xml:"SignedInfo"`
	CanonicalizationMethod Algorithm `xml:"CanonicalizationMethod"`
	SignatureMethod        Algorithm `xml:"SignatureMethod"`
	Reference              Reference
}

// Reference specifies a digest algorithm and digest value, and optionally an identifier of the object being signed, the type of the object, and/or a list of transforms to be applied prior to digesting.
type Reference struct {
	XMLName      xml.Name `xml:"Reference"`
	URI          string   `xml:",attr,omitempty"`
	Transforms   Transforms
	DigestMethod Algorithm `xml:"DigestMethod"`
	DigestValue  string    `xml:"DigestValue"`
}

// Transforms is an optional ordered list of processing steps that were applied to the resource's content before it was digested.
type Transforms struct {
	XMLName   xml.Name    `xml:"Transforms"`
	Transform []Algorithm `xml:"Transform"`
}

// KeyInfo is an optional element that enables the recipient(s) to obtain the key needed to validate the signature.
type KeyInfo struct {
	XMLName  xml.Name `xml:"KeyInfo"`
	X509Data *X509Data
	Children []interface{}
}

// X509Data element within KeyInfo contains one an X509 certificate
type X509Data struct {
	XMLName         xml.Name `xml:"X509Data"`
	X509Certificate string   `xml:"X509Certificate"`
}
