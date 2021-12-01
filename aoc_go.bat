@echo off
docker run --rm -it --name aoc_2021 --mount type=bind,source=D:\Code\Go\AoC2021,target=/app golang bash
cmd /k