# Call Classifier

A program to make manual categorization of tungara frog calls fast.

## How it works

The code uses the Python bindings for [OpenCV](https://opencv-python-tutroals.readthedocs.io/en/latest/py_tutorials/py_tutorials.html) to display images of frog calls to the user. With each image displayed, the user selects a key to classify the type of frog call. The image then gets moves to the folder associated with that category of call.

## Steps

1. Set up the `conda` environment. (You can install `conda` from [here](https://docs.conda.io/en/latest/miniconda.html)). Once `conda` is installed, run `conda config --add channels conda-forge` to make sure you get your packages from conda-forge  then `conda create --name opencv --environment.yaml` to create the environment.
2. `conda activate opencv` to activate the opencv environment.
3. `python classify.py`

The program picks a folder that contains pngs and shows these one at a time to the user. The user selects the key that corresponds with the type of call. That image is then moved to a folder with the name of the call and a new image is shown to the user. Then continues until the folder is exhausted of png files, at which point the program exits.

## Notes on useage

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

The program should be run from, given the above file structure, `some_massive_folder`.

Other filetypes besides `png` files are ignored and will not be moved.

### PDF vs image files

Since the program makes use of OpenCV, it can only display image files like `png`s or `jpg`s. Meghan's images were all in PDF format. To convert all the PDFs to `png`s, I used 

```bash
find . -name *.pdf | gtime parallel -j+0 --eta convert -density 150 {} -quality 90 {.}.png
```

`convert` is an ImageMagick utility. Both ImageMagick and `parallel` can be installed via `brew`. Note that this `convert` takes several seconds for each PDF, so with 39,000 PDFs, this takes awhile.

### For the future

I would like to machine learn these images / calls.