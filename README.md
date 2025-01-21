# Readme for `gamenfo` Package

## 1. Overview
The `gamenfo` package provides a structured representation of game metadata, supporting serialization and deserialization in XML and JSON formats. The package is designed to store detailed information about video games, including release data, multilingual titles and descriptions, and creator permissions. The data model also allows for export into the NFO format for compatibility with media management systems.

## 2. Data Structures

### 2.1 `Game`
The central structure representing a video game.

#### Fields
- **Title** (`[]LangText`): A multilingual array of game titles.
- **Release** (`[]Platform`): Information about platform-specific releases.
- **Genre** (`[]string`): A list of genres associated with the game.
- **Developers** (`[]string`): The development studios or individuals responsible for the game.
- **Publishers** (`[]string`): The entities responsible for publishing the game.
- **Plot** (`[]LangText`): A multilingual array of game plot descriptions.
- **Rating** (`float64`): The aggregate rating of the game.
- **AgeRating** (`[]LangText`): Multilingual descriptions of the game’s age rating.
- **DLCs** (`[]DLC`): A list of downloadable content (DLC) associated with the game.
- **Icon** (`string`): A URL or file path to the game's icon.
- **Thumbnail** (`string`): A URL or file path to the game’s thumbnail image.
- **CreatorInfo** (`*CreatorInfo`): Permissions and guidelines for content creators.
- **CreatedAt** (`time.Time`): The timestamp of when the game entry was created.
- **UpdatedAt** (`time.Time`): The timestamp of when the game entry was last updated.

### 2.2 `Platform`
Represents release information for a specific platform.

#### Fields
- **Platform** (`string`): The platform name (e.g., PC, Xbox).
- **Releases** (`[]CountryDate`): Region-specific release dates.

### 2.3 `CountryDate`
Specifies a release date for a specific country.

#### Fields
- **Country** (`string`): The country code (ISO 3166-1 alpha-2).
- **Date** (`time.Time`): The release date in that country.

### 2.4 `LangText`
Encapsulates multilingual text.

#### Fields
- **Lang** (`string`): The language code (ISO 639-1, optional).
- **Value** (`string`): The text content.

### 2.5 `CreatorInfo`
Defines permissions for content creators regarding the game.

#### Fields
- **IsStreamingAllowed** (`bool`): Whether streaming is permitted.
- **IsVideoOnDemandAllowed** (`bool`): Whether video-on-demand is allowed.
- **MusicUsageAllowed** (`bool`): Whether music usage is allowed.
- **MonetisationAllowed** (`bool`): Whether monetization is permitted.
- **AdditionalInfo** (`string`): Optional additional guidelines or notes.

### 2.6 `DLC`
Represents downloadable content for a game.

#### Fields
- **Title** (`[]LangText`): A multilingual array of DLC titles.
- **ReleaseDate** (`CountryDate`): The release date of the DLC in a specific region.
- **Developers** (`[]string`): The development studios or individuals responsible for the DLC.
- **Publishers** (`[]string`): The entities responsible for publishing the DLC.
- **Plot** (`[]LangText`): A multilingual array of DLC plot descriptions.
- **AgeRating** (`[]LangText`): Multilingual descriptions of the DLC’s age rating.

## 3. Key Methods

### 3.1 `ToNFO`
Generates an NFO file from the `Game` structure.

#### Signature
```go
func (g *Game) ToNFO() (string, error)
```

#### Description
This method serializes the `Game` structure into an indented XML string suitable for use in NFO files.

#### Example Usage
```go
game := Game{
    Title: []LangText{{Lang: "en", Value: "Example Game"}},
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
}
nfo, err := game.ToNFO()
if err != nil {
    log.Fatalf("Failed to generate NFO: %v", err)
}
fmt.Println(nfo)
```

## 4. Supported Serialization Formats

### 4.1 JSON
The `gamenfo` package supports JSON serialization for interoperability with modern web services and APIs.

### 4.2 XML
The XML serialization is designed for compatibility with legacy systems.

## 5. Conformance Requirements

### 5.1 Validation
- All required fields must be populated before serialization.
- `Lang` fields should conform to ISO 639-1 codes if provided.
- `Country` fields in `CountryDate` must conform to ISO 3166-1 alpha-2 codes.

### 5.2 Error Handling
- Serialization methods return errors if mandatory fields are missing or if data cannot be marshaled.
- XML serialization adheres to the 1.0 specification.

## 6. Notes for Developers
- Ensure timestamps (`CreatedAt` and `UpdatedAt`) are in UTC for consistency.
- Use appropriate validation tools to enforce multilingual and regional data integrity.

## 7. Future Extensions
- Add support for additional metadata, such as companion apps.
