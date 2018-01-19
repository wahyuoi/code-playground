package com.gedewahyu.quartz.simple;

import org.quartz.Job;
import org.quartz.JobExecutionContext;
import org.quartz.JobExecutionException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class MyJob implements Job{
  Logger logger = LoggerFactory.getLogger(MyJob.class);
  public void execute(JobExecutionContext jobExecutionContext) throws JobExecutionException {
    logger.info("Hello World!");
  }
}
