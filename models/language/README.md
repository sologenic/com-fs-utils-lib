# Language

The language model is used to set and track the language preference for user or admin accounts, enabling internationalization (i18n) support.

## Enum: `Lang`

Represents supported languages in the system.

### Values

- **`LANG_NOT_USED`** (0): Default/unused value
- **`ENGLISH`** (1): English language
- **`SPANISH`** (2): Spanish language (Español)
- **`KOREAN`** (3): Korean language (한국어)

## Message: `Language`

Represents a user's language preference with metadata about how it was set.

### Fields

- **`Language`** (Lang, required): The selected language from the `Lang` enum
- **`UserConfigured`** (bool, required): Indicates whether the language was explicitly set by the user (true) or is a system default (false)

### Usage

This model allows the system to:
- Store user language preferences
- Distinguish between user-selected languages and system defaults
- Support multi-language interfaces
- Provide localized content based on user preference

### Example Scenarios

- **User sets language**: `Language: ENGLISH, UserConfigured: true`
- **System default**: `Language: ENGLISH, UserConfigured: false`
- **User changes to Spanish**: `Language: SPANISH, UserConfigured: true`

## Use Cases

- User account settings
- Admin panel localization
- Email template selection
- UI language switching
- Content localization

## Notes

- The `UserConfigured` flag helps distinguish explicit user choices from defaults
- This enables features like "detect language" vs "user selected language"
- Additional languages can be added to the `Lang` enum as needed
- When `UserConfigured` is false, the system may use browser/device language detection
