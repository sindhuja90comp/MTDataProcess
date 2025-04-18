package com.example;

import java.io.IOException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class DataProcessorJava {
    public static void main(String[] args) throws InterruptedException, IOException {
        int numWorkers = 5;
        int numTasks = 20;
        SharedQueue taskQueue = new SharedQueue();
        ResultSaver resultSaver = new ResultSaver("results_java.txt");
        ExecutorService executorService = Executors.newFixedThreadPool(numWorkers);

        // Add tasks to the queue
        for (int i = 0; i < numTasks; i++) {
            taskQueue.addTask(new Task(String.valueOf(i), "Data-" + i));
        }

        // Start worker threads
        for (int i = 0; i < numWorkers; i++) {
            executorService.execute(new Worker(String.valueOf(i + 1), taskQueue, resultSaver));
        }

        // Signal workers to stop when no more tasks
        executorService.shutdown();
        executorService.awaitTermination(5, TimeUnit.SECONDS);

        // Close the result saver
        resultSaver.close();
        System.out.println("All tasks submitted and workers finished.");
    }
}