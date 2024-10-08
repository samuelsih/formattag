## Formattag
This is forked project from [momaek/formattag](https://github.com/momaek/formattag).

The tool is used to align golang struct's tags.

eg.:

Before
```golang
// TestStruct this is a test struct
type TestStruct struct {
	ID            string `json:"id" xml:"id"`
	IfNotModified string `json:"if_not_modified" xml:"if_not_modified"`
	Name          string `json:"name" xml:"name"`

	ThisIsAStructWodeTianNa struct {
		FieldFromThisIsAStructWodeTianNa string `json:"field_from_this_is_a_struct_wode_tian_na" xml:"field_from_this_is_a_struct_wode_tian_na"`
		TianName                         string `json:"tian_name" xml:"tian_name"`
	} `json:"this_is_a_struct_wode_tian_na" xml:"this_is_a_struct_wode_tian_na"`

	T    time.Time `json:"t" xml:"t"`
	Fset Fset      `json:"fset" xml:"fset"`
}

type Fset struct{}

```

After
```golang
// TestStruct this is a test struct
type TestStruct struct {
	ID            string `json:"id"              xml:"id"`
	IfNotModified string `json:"if_not_modified" xml:"if_not_modified"`
	Name          string `json:"name"            xml:"name"`

	ThisIsAStructWodeTianNa struct {
		FieldFromThisIsAStructWodeTianNa string `json:"field_from_this_is_a_struct_wode_tian_na" xml:"field_from_this_is_a_struct_wode_tian_na"`
		TianName                         string `json:"tian_name"                                xml:"tian_name"`
	} `json:"this_is_a_struct_wode_tian_na" xml:"this_is_a_struct_wode_tian_na"`

	T    time.Time `json:"t"    xml:"t"`
	Fset Fset      `json:"fset" xml:"fset"`
}

type Fset struct{}
```

### Installation

####  Using go install

```
go install github.com/samuelsih/formattag@latest
```

### Usage

```
formattag .
```