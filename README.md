# My Simple Password Manager

Is is a simple password manager for your terminal with passwords decoding.

## Args

> **add** "Key" "Value"  // Add new password to db
>
> ```sh
>~ password-manager -add "key" "password"
>```
>
> **get** "Key"  // Get password by the key
>
> ```sh
>~ password-manager -get "key"
>```
>
> **keys**  // Get table of all keys with passwords
>
> ```sh
>~ password-manager -keys
>```

## Configuration

In the configuration file you can change database path and theme of the output data.

```json
{
    "theme": "StyleColoredDefault",
    "database_path": "database/pm.db"
}
```

## Full list of themes

> - StyleBold
> - StyleDefault
> - StyleDouble
> - StyleLight
> - StyleRounded
> - StyleColoredBlackOnBlueWhite
> - StyleColoredBlackOnGreenWhite
> - StyleColoredBlackOnCyanWhite
> - StyleColoredBlackOnMagentaWhite
> - StyleColoredBlackOnRedWhite
> - StyleColoredBlackOnYellowWhite
> - StyleColoredBlueWhiteOnBlack
> - StyleColoredGreenWhiteOnBlack
> - StyleColoredMagentaWhiteOnBlack
> - StyleColoredRedWhiteOnBlack
> - StyleColoredYellowWhiteOnBlack
> - StyleColoredBright
> - StyleColoredDark <-- Default
