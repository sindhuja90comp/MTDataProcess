package com.example;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class ResultSaver {
    private List<String> results = new ArrayList<>();
    private Lock lock = new ReentrantLock();
    private FileWriter fileWriter;
    private String filename;

    public ResultSaver(String filename) throws IOException {
        // Construct the file path using the provided relative path "java/results"
        String basePath = new File("").getAbsolutePath();
        System.out.println("hello world" + basePath);
        this.filename = basePath + File.separator + results + File.separator + filename;
        System.out.println(this.filename);

        // Ensure the directory exists
        File file = new File(this.filename);
        file.getParentFile().mkdirs();
        this.fileWriter = new FileWriter(file);
    }

    public void saveResult(String result) {
        lock.lock();
        try {
            results.add(result);
            fileWriter.write(result + "\n");
        } catch (IOException e) {
            System.err.println("Error writing to file: " + e.getMessage());
        } finally {
            lock.unlock();
        }
    }

    public void close() throws IOException {
        fileWriter.close();
    }
}
