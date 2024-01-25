import os
import json
import re
import sys

# flag用于标识节点1还是其余所有节点，false标
flag = sys.argv[0]

PATH = "./docker_data/node"

enode = ""
if flag:
    # 获取节点1的enode
    with open("enode.txt", "r") as f:
        enode_res = f.readlines()[0]
    if "discport" in enode_res:
        enode = enode_res.split("?")[0][1:]
    else:
        enode = enode_res[1:-2]




# 遍历更改其余节点启动文件
for i in range(1,5):
    # 获取各节点addresss
    workdir = os.path.join(PATH+str(i), "keystore")
    file_names = os.listdir(workdir)
    if len(file_names) == 1:
        file_path = os.path.join(workdir, file_names[0])
        print("文件路径：", file_path)
    else:
        print("文件夹中不止一个文件或者文件夹为空。")
    with open(file_path, "r") as f:
        cnts = json.load(f)

    address = "0x" + cnts["address"]

    if i == 1:
        # 更改genesis.json
        genesis_path = "docker_data/genesis.json"
        with open(genesis_path, "r") as f1, open("%s.bak" % genesis_path, "w") as f2:
            for idx, line in enumerate(f1):
                if idx == 25 or idx == 33:
                    line = re.sub(re.findall(r'"(.*?)"', line)[0], address, line)
                    print(line)
                f2.write(line)
            os.remove(genesis_path)
            os.rename("%s.bak" % genesis_path, genesis_path)

    # 改启动文件
    sh_path = "docker_data/start_node"+str(i)+".sh"
    with open(sh_path, "r") as f1, open("%s.bak" % sh_path, "w") as f2:
        for line in f1:
            if "unlock" in line:
                line = re.sub(re.findall(r'--unlock\s+(\w+)', line)[0], address, line)
                # print(line)
            elif "admin" in line:
                if flag:
                    print(enode)
                    line = re.sub(re.findall(r'"(.*?)"', line)[0], enode, line)
            f2.write(line)
        os.remove(sh_path)
        os.rename("%s.bak" % sh_path, sh_path)
