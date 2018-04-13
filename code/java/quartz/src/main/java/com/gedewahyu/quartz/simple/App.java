package com.gedewahyu.quartz.simple;

import static org.quartz.JobBuilder.newJob;
import static org.quartz.SimpleScheduleBuilder.simpleSchedule;
import static org.quartz.TriggerBuilder.newTrigger;

import org.quartz.JobDetail;
import org.quartz.Scheduler;
import org.quartz.SchedulerException;
import org.quartz.Trigger;
import org.quartz.impl.StdSchedulerFactory;

public class App {

  public static void main(String[] args) {
    new App().run();
  }

  private void run() {
    try {
      Scheduler scheduler = StdSchedulerFactory.getDefaultScheduler();

      JobDetail job = newJob(MyJob.class)
          .withIdentity("job1", "group1")
          .build();

      Trigger trigger = newTrigger()
          .withIdentity("trigger1", "group1")
          .startNow()
          .withSchedule(simpleSchedule()
              .withIntervalInSeconds(5)
              .repeatForever())
          .build();

      scheduler.scheduleJob(job, trigger);
      scheduler.start();
    } catch (SchedulerException e) {
      e.printStackTrace();
    }
  }
}
