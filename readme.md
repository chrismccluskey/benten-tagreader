# Project Benten: Tag Reader

#### Current Status: Alpha

## Purpose

Provides search capabilities to the rest of project benten, to be used by agents which allow remote management of music libraries.

## What it does today

Supports only mp3 files at this time, though broader support is planned for the near future. Scans a directory of files, read all id3 frames and return which files matched a specific frame key/value pair or print all frames.

## Example Usage

Print file name and all mp3 frames for all files in the current directory
```
./benten-tagreader --print-frames
```

Find all files in a directory that contain a frame which matches a specific value. For example, the TIT1 "Grouping" frame which contains the value "Instrumental"
```
./benten-tagreader --match-frame=TIT1 --match-text=Instrumental
```