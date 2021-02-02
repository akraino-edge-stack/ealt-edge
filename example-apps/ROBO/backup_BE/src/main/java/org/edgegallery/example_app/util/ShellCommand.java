package org.edgegallery.example_app.util;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.List;
import org.springframework.stereotype.Service;

@Service
public class ShellCommand {

    public String executeCommand(String command) {

        StringBuffer output = new StringBuffer();

        Process p;
        try {
            p = Runtime.getRuntime().exec(command);
            p.waitFor();
            BufferedReader reader =
                    new BufferedReader(new InputStreamReader(p.getInputStream()));

            String line = "";
            while ((line = reader.readLine())!= null) {
                output.append(line + "\n");
            }

        } catch (Exception e) {
            e.printStackTrace();
        }

        return output.toString();

    }

    //parse velero cmd and get details
    public String parseResult(String msg){
        List<String> itemsList = new ArrayList<String>();

        /*
        if (msg == null || msg.equals(""))
            return itemsList;

        matcher = pattern.matcher(msg);
        while (matcher.find()) {
            ipList.add(matcher.group(0));
        }
        return ipList;
*/
        return "success";
    }
}