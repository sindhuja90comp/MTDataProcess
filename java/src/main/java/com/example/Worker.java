package com.example;

public class Worker implements Runnable {
    private SharedQueue taskQueue;
    private ResultSaver resultSaver;
    private String workerId;

    public Worker(String id, SharedQueue queue, ResultSaver saver) {
        this.workerId = id;
        this.taskQueue = queue;
        this.resultSaver = saver;
        System.out.println("Worker " + workerId + " started.");
    }

    @Override
    public void run() {
        while (true) {
            try {
                Task task = taskQueue.getTask();
                if (task == null) {
                    System.out.println("Worker " + workerId + " finished (no more tasks).");
                    break;
                }
                String result = task.process();
                resultSaver.saveResult(result);
                System.out.println("Worker " + workerId + " processed task " + task.getId());
            } catch (InterruptedException e) {
                System.err.println("Worker " + workerId + " interrupted: " + e.getMessage());
                Thread.currentThread().interrupt();
                break;
            }
        }
    }
}