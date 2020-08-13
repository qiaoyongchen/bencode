# Bencode 介绍

BEncoding是一种编码方式，比如种子文件就是就是采用这种编码方式。

Bencode有4种类型数据:

## 1. String

"12345" => 5:12345

## 2. Int

12345 => i12345e

## 3. List

List<"abced", 12345> => l5:abcdei12345ee

## 4. dictionary

Dictionary<{"abced":"abced"},{"abc":123}> => d5:abced5:abced3:abci23ee

# 使用方式

