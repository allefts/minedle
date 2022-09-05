import os
#  mongoimport .\acacia_fence_gate.json -d MinecraftItems -c Items --uri 'mongodb+srv://allefts:sa680108@minedledb.vepcvkm.mongodb.net/?retryWrites=true&w=majority'

directory = './'
 
# iterate over files in
# that directory
for filename in os.listdir(directory):
    f = os.path.join(directory, filename)
    # checking if it is a file
    if os.path.isfile(f):
        os.system("mongoimport " + f + " -d MinecraftItems -c Items --uri mongodb+srv://allefts:sa680108@minedledb.vepcvkm.mongodb.net/?retryWrites=true&w=majority --drop")