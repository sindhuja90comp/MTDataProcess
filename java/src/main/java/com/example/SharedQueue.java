package com.example;

import java.util.LinkedList;
import java.util.Queue;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class SharedQueue {
    private Queue<Task> queue = new LinkedList<>();
    private Lock lock = new ReentrantLock();

    public void addTask(Task task) {
        lock.lock();
        try {
            queue.offer(task);
            notifyAll(); // Notify waiting threads that a task has been added
        } finally {
            lock.unlock();
        }
    }

    public Task getTask() throws InterruptedException {
        lock.lock();
        try {
            while (queue.isEmpty()) {
                wait(); // Use wait/notify with the lock
            }
            return queue.poll();
        } finally {
            lock.unlock();
        }
    }

    public  boolean isEmpty() { //Removed synchronized
        lock.lock();
        try{
            return queue.isEmpty();
        }finally{
            lock.unlock();
        }

    }
}
