package com.example;

import java.util.Random;

public class Task {
    private String id;
    private String data;

    public Task(String id, String data) {
        this.id = id;
        this.data = data;
    }

    public String getId() {
        return id;
    }

    public String getData() {
        return data;
    }

    public String process() {
        // Simulate processing time
        try {
            Thread.sleep(new Random().nextInt(500));
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        return "Processed: " + data + " (Task ID: " + id + ")";
    }
}