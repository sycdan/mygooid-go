# MyGooid

This is my first Golang project.

The goal is to create a unique identifier for a person, which combines both public and private information about them. This could be used to prove their identity in an authentication flow.

## Card generation

We randomize the order of 72 characters: 26 uppercase, 26 lowercase, 10 numbers, 10 symbols. This randomized character set is rendered onto a 9x9 grid, leaving space inside each internal 3x3 grid for an "anchor".

Anchors are composed of letters and numbers that can appear in the person's memorable secret.

## Running

```bash
go run . "Full legal name" "Memorable secret"
```
