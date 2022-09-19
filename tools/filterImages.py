import os

imageDirectory = "../recipe-images/"
recipeDirectory = "../recipes/"

recipesUsed = []
imagesUsed = []
noImages = []

for filename in os.listdir(recipeDirectory):
    recipesUsed.append(filename.removesuffix(".json"))

for filename in os.listdir(imageDirectory):
    imagesUsed.append(filename.removesuffix(".png"))

allUsed = set(recipesUsed) & set(imagesUsed)

#for filename in os.listdir(recipeDirectory):
#    if filename.removesuffix(".json") not in imagesUsed:
#        noImages.append(filename.removesuffix(".json"))
#        print(filename)

for filename in os.listdir(imageDirectory):
    if filename.removesuffix(".png") not in recipesUsed:
        os.remove(imageDirectory + filename)

