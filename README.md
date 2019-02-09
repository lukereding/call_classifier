# Call Classifier

A program to make manual categorization of tungara frog calls fast.

## How it works

The code uses [GoCV](https://gocv.io/) to display images of frog calls to the user. With each image displayed, the user selects a key to classify the type of frog call. The image then gets moves to the folder associated with that category of call.

## Installation

The program is available as an executable binary compiled on macOS, which means all you need to do is download the binary, make it executable `chmod +x call_classifer`, then execute it with `./call_classifer` or `bash call_classifer`. However, you must have OpenCV v.4 installed on your machine. [This](https://gocv.io/getting-started/macos/) page outlines the steps, but `brew install hybridgroup/tools/opencv` should be all you need, assuming you've installed [homebrew](brew.sh).

## Notes

The program assumes these are the types are calls you care about. Type of call on the left, keystroke to press to denote that call on the right:

| call type | keystroke | 
|---|---|
| whine | 0 |
| whine with one chuck | 1 |
| whine with two chucks | 2 |
| ... | ... |
| whine with eight chucks | 8 |
| mew | 9 |
| unknown / other | u | 

These can be changed easily, or modified so that this table is read in via a csv file or something similar.

The program will create a folder for each of these call types in the top-level directory from which the program is called. The program also assumes that the current file structure for your images, which looks something like this

```
some_massive_folder
├── folderWithImagesFromOneTrial
│   ├── image1.png
│   ├── image2.png
│   ├── image3.png
│   ├── ...
```

Other filetypes besides `png` files are ignored and will not be moved.

Each time you call the program, it will find a single folder containing `png`s. It will exit when done with that folder. Call the program again to start on a new folder.

### PDF vs image files

Since the program makes use of [GoCV](https://gocv.io/), it can only display image files like `png`s or `jpg`s. Meghan's images were all in PDF format. To convert all the PDFs to `png`s, I used 

```bash
find . -name *.pdf | gtime parallel -j+0 --eta convert -density 150 {} -quality 90 {.}.png
```

`convert` is an ImageMagick utility. Both ImageMagick and `parallel` can be installed via `brew`. Note that this `convert` takes several seconds for each PDF, so with 39,000 PDFs, this takes awhile.

### For the future

I would like to machine learn these images / calls.